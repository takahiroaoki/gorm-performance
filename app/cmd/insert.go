package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/takahiroaoki/gorm-performance/app/config"
	"github.com/takahiroaoki/gorm-performance/app/dao"
	"github.com/takahiroaoki/gorm-performance/app/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewInsertCmd() *cobra.Command {
	var rep int
	var sdt bool
	var ps bool

	insertCmd := &cobra.Command{
		Use:   "insert",
		Short: "Exec command for insert test",
		Long:  "Exec command for insert test",
		RunE: func(cmd *cobra.Command, args []string) error {

			// Prepare db client
			gormConfig := &gorm.Config{
				SkipDefaultTransaction: sdt,
				PrepareStmt:            ps,
			}
			db, err := gorm.Open(
				mysql.Open(config.GetDataSourceName()),
				gormConfig,
			)
			if err != nil {
				log.Fatalf(fmt.Sprintf("Failed to get DB connection. Error: %v", err))
			}
			defer closeDB(db)

			dao := dao.NewRecordDao()
			rec := entity.Record{
				Test: "test",
			}

			// Measurement
			err = db.Transaction(func(tx *gorm.DB) error {
				start := time.Now()
				for i := 0; i < rep; i++ {
					_, err = dao.InsertRecord(tx, rec)
					if err != nil {
						return err
					}
				}
				time := time.Since(start)
				fmt.Printf("repeat: %v\n", rep)
				fmt.Printf("time: %v\n", time)

				return nil
			})

			if err != nil {
				log.Fatalf("[ERROR] %v", err)
			}

			return nil
		},
	}

	insertCmd.Flags().IntVarP(&rep, "repeat", "r", 100, "Repeat count")
	insertCmd.Flags().BoolVarP(&sdt, "skip-default-transaction", "s", false, "SkipDefaultTransaction of gorm.Config")
	insertCmd.Flags().BoolVarP(&sdt, "prepare-statement", "p", false, "PrepareStatement of gorm.Config")
	return insertCmd
}

func closeDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln("Failed to close db connection")
		return
	}
	if err := sqlDB.Close(); err != nil {
		log.Fatalln("Failed to close db connection")
		return
	}
}
