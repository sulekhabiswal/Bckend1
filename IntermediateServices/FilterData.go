package IntermediateServices

import "fmt"

//Filters Customer Data
func FilterCustomerData(data map[string]interface{}, bookMark string) (map[string]interface{}, error) {
	// Define null values
	nullArr := map[interface{}]bool{nil: true, "null": true, "": true}

	// Function to replace null values with nil
	getColumnValue := func(column interface{}) interface{} {
		if _, exists := nullArr[column]; exists {
			return nil
		}
		return column
	}

	// Filter customer data
	filteredData := map[string]interface{}{
		"user_id":    getColumnValue(data["column1"]),
		"user_name":  getColumnValue(data["column2"]),
		"first_name": getColumnValue(data["column3"]),
		// "middle_name":     getColumnValue(data["column4"]),
		"last_name":       getColumnValue(data["column5"]),
		"user_role":       getColumnValue(data["column7"]),
		"address":         getColumnValue(data["column8"]),
		"city":            getColumnValue(data["column9"]),
		"state":           getColumnValue(data["column10"]),
		"pincode":         getColumnValue(data["column11"]),
		"self_onboard":    getColumnValue(data["column12"]),
		"email":           getColumnValue(data["column13"]),
		"mobile_number":   getColumnValue(data["column14"]),
		"approver_role":   getColumnValue(data["column15"]),
		"approver_id":     getColumnValue(data["column16"]),
		"approved_by":     getColumnValue(data["column17"]),
		"created_by":      getColumnValue(data["column18"]),
		"is_active":       getColumnValue(data["column25"]),
		"admin_name":      getColumnValue(data["column26"]),
		"creator_id":      getColumnValue(data["column27"]),
		"creator_role":    getColumnValue(data["column28"]),
		"created_date":    getColumnValue(data["column29"]),
		"updated_date":    getColumnValue(data["column30"]),
		"address1":        getColumnValue(data["column31"]),
		"admin_role":      getColumnValue(data["column32"]),
		"channel":         getColumnValue(data["column33"]),
		"country":         getColumnValue(data["column34"]),
		"is_child":        getColumnValue(data["column35"]),
		"is_block":        getColumnValue(data["column37"]),
		"block_remarks":   getColumnValue(data["column38"]),
		"block_timestamp": getColumnValue(data["column39"]),
		"admin_id":        getColumnValue(data["column41"]),
		"kyc_status":      getColumnValue(data["column42"]),
		"kyc_type":        getColumnValue(data["column43"]),
		"is_fraud":        getColumnValue(data["column44"]),
		"fraud_remarks":   getColumnValue(data["column45"]),
		"fraud_timestamp": getColumnValue(data["column46"]),
		"remarks":         getColumnValue(data["column47"]),
		"param_a":         getColumnValue(data["column48"]),
		"param_b":         getColumnValue(data["column49"]),
		"param_c":         getColumnValue(data["column50"]),
		"kyc_number":      getColumnValue(data["column54"]),
		"gender":          getColumnValue(data["column64"]),
		"date_of_birth":   getColumnValue(data["column65"]),
		"latlong":         getColumnValue(data["column72"]),
	}

	fmt.Println("INFO:", bookMark, "Filtered Customer Data:", filteredData)

	return filteredData, nil
}

// Filters Card Data
func FilterCardData(data map[string]interface{}, bookMark string) (map[string]interface{}, error) {
	// Define null values
	nullArr := map[interface{}]bool{nil: true, "null": true, "": true}

	// Function to replace null values with nil
	getColumnValue := func(column interface{}) interface{} {
		if _, exists := nullArr[column]; exists {
			return nil
		}
		return column
	}

	// Filter card data
	filteredData := map[string]interface{}{
		"card_refnumber": getColumnValue(data["column1"]),
		"customer_id":    getColumnValue(data["column2"]),
		"created_date":   getColumnValue(data["column29"]),
		"updated_date":   getColumnValue(data["column30"]),

		// "card_refnumber":         getColumnValue(data["column1"]),
		// "customer_id":            getColumnValue(data["column2"]),
		// "atm_limit":              getColumnValue(data["column3"]),
		// "bin_number":             getColumnValue(data["column4"]),
		// "parent_customer_id":     getColumnValue(data["column5"]),
		// "ecommerce_limit":        getColumnValue(data["column7"]),
		// "is_ecommerce_allowed":   getColumnValue(data["column8"]),
		// "pos_limit":              getColumnValue(data["column9"]),
		// "is_pos_allowed":         getColumnValue(data["column10"]),
		// "is_atmallowed":          getColumnValue(data["column11"]),
		// "mobile_number":          getColumnValue(data["column14"]),
		// "card_brand":             getColumnValue(data["column15"]),
		// "card_type":              getColumnValue(data["column16"]),
		// "card_limit":             getColumnValue(data["column17"]),
		// "encrypted_card_number":  getColumnValue(data["column18"]),
		// "last_fourdigit":         getColumnValue(data["column19"]),
		// "jit_allowed":            getColumnValue(data["column20"]),
		// "wallet_status":          getColumnValue(data["column21"]),
		// "wallet_id":              getColumnValue(data["column22"]),
		// "service_code":           getColumnValue(data["column23"]),
		// "expiry_date":            getColumnValue(data["column24"]),
		// "is_active":              getColumnValue(data["column25"]),
		// "pinmailer_flag":         getColumnValue(data["column26"]),
		// "pinoffset":              getColumnValue(data["column27"]),
		// "pin_try_count":          getColumnValue(data["column28"]),
		// "created_date":           getColumnValue(data["column29"]),
		// "updated_date":           getColumnValue(data["column30"]),
		// "product_id":             getColumnValue(data["column31"]),
		// "product_name":           getColumnValue(data["column32"]),
		// "product_type":           getColumnValue(data["column33"]),
		// "product_category":       getColumnValue(data["column34"]),
		// "is_child":               getColumnValue(data["column35"]),
		// "is_block":               getColumnValue(data["column37"]),
		// "block_remarks":          getColumnValue(data["column38"]),
		// "block_timestamp":        getColumnValue(data["column39"]),
		// "is_fraud":               getColumnValue(data["column44"]),
		// "fraud_remarks":          getColumnValue(data["column45"]),
		// "fraud_timestamp":        getColumnValue(data["column46"]),
		// "name_on_card":           getColumnValue(data["column47"]),
		// "param_a":                getColumnValue(data["column48"]),
		// "param_b":                getColumnValue(data["column49"]),
		// "param_c":                getColumnValue(data["column50"]),
		// "shipping_addressline1":  getColumnValue(data["column54"]),
		// "shipping_addressline2":  getColumnValue(data["column55"]),
		// "shipping_city":          getColumnValue(data["column56"]),
		// "shipping_country":       getColumnValue(data["column57"]),
		// "shipping_pincode":       getColumnValue(data["column58"]),
		// "shipping_state":         getColumnValue(data["column59"]),
		// "is_close":               getColumnValue(data["column60"]),
		// "is_embossa":             getColumnValue(data["column61"]),
		// "is_green_pin":           getColumnValue(data["column62"]),
		// "is_hotlist":             getColumnValue(data["column63"]),
		// "is_lost":                getColumnValue(data["column64"]),
		// "is_parallel_add_on":     getColumnValue(data["column65"]),
		// "is_physical_pin":        getColumnValue(data["column66"]),
		// "is_si":                  getColumnValue(data["column67"]),
		// "is_stolen":              getColumnValue(data["column68"]),
		// "is_suspended":           getColumnValue(data["column69"]),
		// "is_virtual":             getColumnValue(data["column70"]),
		// "is_damage":              getColumnValue(data["column71"]),
		// "damage_timestamp":       getColumnValue(data["column72"]),
		// "lost_timestamp":         getColumnValue(data["column73"]),
		// "stolen_timestamp":       getColumnValue(data["column74"]),
		// "de_active":              getColumnValue(data["column75"]),
		// "deactive_timestamp":     getColumnValue(data["column76"]),
		// "hotlist_timestamp":      getColumnValue(data["column77"]),
		// "is_contactless_allowed": getColumnValue(data["column78"]),
		// "per_transaction_max_limit_in_contactless": getColumnValue(data["column79"]),
	}

	fmt.Println("INFO:", bookMark, "Filtered Card Data:", filteredData)

	return filteredData, nil
}

//Filter Card_ACS

func FilterCardAcsData(data map[string]interface{}, bookMark string) (map[string]interface{}, error) {
	// Define null values
	nullArr := map[interface{}]bool{nil: true, "null": true, "": true}

	// Function to replace null values with nil
	getColumnValue := func(column interface{}) interface{} {
		if _, exists := nullArr[column]; exists {
			return nil
		}
		return column
	}

	filteredData := map[string]interface{}{
		"encrypted_card_number": getColumnValue(data["column1"]),
		"card_refnumber":        getColumnValue(data["column2"]),
		"mobile_number":         getColumnValue(data["column14"]),
		"expiry_date":           getColumnValue(data["column24"]),
		"is_active":             getColumnValue(data["column25"]),
		"created_date":          getColumnValue(data["column29"]),
		"updated_date":          getColumnValue(data["column30"]),
		"is_si_enable":          getColumnValue(data["column48"]),
		"is_block":              getColumnValue(data["column37"]),
		"is_fraud":              getColumnValue(data["column44"]),
		"is_hotlist":            getColumnValue(data["column63"]),
		"is_lost":               getColumnValue(data["column64"]),
		"is_stolen":             getColumnValue(data["column68"]),
	}

	fmt.Println("INFO:", bookMark, "Filtered Card Data:", filteredData)

	return filteredData, nil
}
