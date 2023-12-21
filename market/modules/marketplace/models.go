package marketplace

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// "github.com/wantedly/gorm-zap"
	// "go.uber.org/zap"

	"argomarket/market/modules/util"
	"argomarket/market/modules/settings"
)

var (
	database *gorm.DB
)

func SyncModels() {
	database.AutoMigrate(
		&Notification{},
		&Category{},
		// &Advertising{},
		// &APISession{},
		// &BitcoinCashTransaction{},
		// &BitcoinTransaction{},
		// &CityMetroStation{},
		// &City{},
		// &Country{},
		// &DepositHistory{},
		// &Deposit{},
		// &DisputeClaim{},
		// &Dispute{},
		// &EthereumTransaction{},
		// &FeedItem{},
		// &ItemCategory{},
		// &Item{},
		// &MessageboardSection{},
		// &Message{},
		&PackagePrice{},
		&Package{},
		&ShippingMethod{},
		// &PaymentReceipt{},
		// &RatingReview{},
		// &ReferralPayment{},
		// &Reservation{},
		// &ShippingOption{},
		// &ShippingStatus{},
		// &StoreUser{},
		// &Store{},
		// &SupportTicketStatus{},
		// &SupportTicket{},
		// &ThreadPerusalStatus{},
		// &TransactionStatus{},
		// &Transaction{},
		// &UserBitcoinCashWalletAction{},
		// &UserBitcoinCashWalletBalance{},
		// &UserBitcoinCashWallet{},
		// &UserBitcoinWalletAction{},
		// &UserBitcoinWalletBalance{},
		// &UserBitcoinWallet{},
		// &UserEthereumWalletAction{},
		// &UserEthereumWalletBalance{},
		// &UserEthereumWallet{},
		// &UserSettingsHistory{},
		// &StoreWarning{},
		&User{},
	)
}

func SeedCategories() error {
	drugs_chemicals := Category{Uuid: util.GenerateUuid(), Name: "Drugs & Chemicals", ParentCategoryUuid: ""}
	counterifeit_items := Category{Uuid: util.GenerateUuid(), Name: "Counterifeit Items", ParentCategoryUuid: ""}
	jewels_gold := Category{Uuid: util.GenerateUuid(), Name: "Jewels & Gold", ParentCategoryUuid: ""}
	carded_items := Category{Uuid: util.GenerateUuid(), Name: "Carded Items", ParentCategoryUuid: ""}
	services := Category{Uuid: util.GenerateUuid(), Name: "Services", ParentCategoryUuid: ""}
	others := Category{Uuid: util.GenerateUuid(), Name: "Other Listings", ParentCategoryUuid: ""}

	fmt.Println("loading seed categories....")
	categories := []Category{drugs_chemicals, counterifeit_items, jewels_gold, carded_items, services, others}

	for _, category := range categories {
		err := database.Create(&category).Error
		if err != nil {
			// Handle the error
			fmt.Println("Error creating category:", err)
		}
	}

	benzos := Category{Uuid: util.GenerateUuid(), Name: "Benzos", ParentCategoryUuid: drugs_chemicals.Uuid}
	canabis_hashish := Category{Uuid: util.GenerateUuid(), Name: "Canabis & Hashish", ParentCategoryUuid: drugs_chemicals.Uuid}
	rc := Category{Uuid: util.GenerateUuid(), Name: "RC", ParentCategoryUuid: drugs_chemicals.Uuid}
	other := Category{Uuid: util.GenerateUuid(), Name: "Other", ParentCategoryUuid: drugs_chemicals.Uuid}

	drugs_chemicals_c := []Category{benzos, canabis_hashish, rc, other}

	for _, category := range drugs_chemicals_c {
		err := database.Create(&category).Error
		if err != nil {
			// Handle the error
			fmt.Println("Error creating category:", err)
		}
	}

	benzos1 := Category{Uuid: util.GenerateUuid(), Name: "Benzos1", ParentCategoryUuid: benzos.Uuid}
	benzos2 := Category{Uuid: util.GenerateUuid(), Name: "Benzos2", ParentCategoryUuid: benzos.Uuid}
	benzos3 := Category{Uuid: util.GenerateUuid(), Name: "Benzos3", ParentCategoryUuid: benzos.Uuid}
	benzos4 := Category{Uuid: util.GenerateUuid(), Name: "Benzos4", ParentCategoryUuid: benzos.Uuid}

	benzos_c := []Category{benzos1, benzos2, benzos3, benzos4}

	for _, category := range benzos_c {
		err := database.Create(&category).Error
		if err != nil {
			// Handle the error
			fmt.Println("Error creating category:", err)
		}
	}

	return nil
}

func SyncDatabaseViews() {

	// drop all views and triggers

	database.Exec(`
		CREATE OR REPLACE FUNCTION strip_all_triggers() RETURNS text AS $$ DECLARE
	    triggNameRecord RECORD;
	    triggTableRecord RECORD;
	BEGIN
	    FOR triggNameRecord IN select distinct(trigger_name) from information_schema.triggers where trigger_schema = 'public' LOOP
	        FOR triggTableRecord IN SELECT distinct(event_object_table) from information_schema.triggers where trigger_name = triggNameRecord.trigger_name LOOP
	            RAISE NOTICE 'Dropping trigger: % on table: %', triggNameRecord.trigger_name, triggTableRecord.event_object_table;
	            EXECUTE 'DROP TRIGGER ' || triggNameRecord.trigger_name || ' ON ' || triggTableRecord.event_object_table || ';';
	        END LOOP;
	    END LOOP;

	    RETURN 'done';
	END;
	$$ LANGUAGE plpgsql SECURITY DEFINER;
	`)

	database.Exec(`
		select strip_all_triggers();
	`)

	database.Exec(`
	SELECT 
	'DROP VIEW ' || table_name || ';'
	FROM information_schema.views
	WHERE table_schema NOT IN ('pg_catalog', 'information_schema')
	AND table_name !~ '^pg_';
	`)

	// wallets & balances
	// setupUserBitcoinBalanceViews()
	// setupUserBitcoinCashBalanceViews()
	// setupUserEthereumBalanceViews()

	// messageboard & messages
	// setupThreadsViews()
	// setupMessageboardThreadsViews()
	// setupPrivateThreadsFunctions()
	// setupVendorVerificationThreadsFunctions()
	// setupMessageboardCategoriesViews()

	// transcations
	// setupTransactionStatusesView()

	// users
	setupUserViews()
	// setupVendorTxStatsViews()
	// setupItemTxStatsViews()

	// items & packages, categories
	// setupCategoriesViews()
	// setupPackagesView()
	// setupSerpItemsView()
	// setupFrontPageItemsViews()

	// tickets
	// setupSupportTicketViews()

	// advertisings
	// setupAdvertisingViews()
}

func init() {
	var err error

	database, err = gorm.Open("postgres", "postgres://localhost:5432/go_t?user=postgres&password=password&sslmode=disable")
	if err != nil {
		panic(err)
	}
	database.DB().SetMaxOpenConns(30)
	database.DB().Ping()

	if settings.GetSettings().Debug {

		// logger, err := zap.NewProduction()
		// if err != nil {
		// 	panic(err)
		// }

		database.LogMode(true)
		// database.SetLogger(gormzap.New(logger))
	}
}
