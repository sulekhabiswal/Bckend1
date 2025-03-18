package IntermediateServices

import "fmt"


func FilterDataBigquery(data map[string]interface{}, bookMark string) (map[string]interface{}, error) {
	nullArr := map[interface{}]bool{nil: true, "null": true, "": true}

	getColumnValue := func(column interface{}) interface{} {
		if _, exists := nullArr[column]; exists {
			return nil
		}
		return column
	}

	filteredData := map[string]interface{}{
		"id":               getColumnValue(data["id"]),
		"user_id":          getColumnValue(data["user_id"]),
		"acquirer_id":      getColumnValue(data["acquirer_id"]),
		"admin_id":         getColumnValue(data["admin_id"]),
		"admin_name":       getColumnValue(data["admin_name"]),
		"admin_role":       getColumnValue(data["admin_role"]),
		"admin_wallet_id":  getColumnValue(data["admin_wallet_id"]),
		"amount":           getColumnValue(data["amount"]),
		"approved_by":      getColumnValue(data["approved_by"]),
		"approved_date":    getColumnValue(data["approved_date"]),
		"approver_id":      getColumnValue(data["approver_id"]),
		"approver_role":    getColumnValue(data["approver_role"]),
		"atm_indicator":    getColumnValue(data["atm_indicator"]),
		"auth_code":        getColumnValue(data["auth_code"]),
		"card_ref_id":      getColumnValue(data["card_ref_id"]),
		"card_type":        getColumnValue(data["card_type"]),
		"created_date":     getColumnValue(data["created_date"]),
		"debit_type":       getColumnValue(data["debit_type"]),
		"device_details":   getColumnValue(data["device_details"]),
		"is_approve":       getColumnValue(data["is_approve"]),
		"is_requested":     getColumnValue(data["is_requested"]),
		"location":         getColumnValue(data["location"]),
		"mcc":             getColumnValue(data["mcc"]),
		"network_id":       getColumnValue(data["network_id"]),
		"operation_mode":   getColumnValue(data["operation_mode"]),
		"operation_performed": getColumnValue(data["operation_performed"]),
		"receipt_link":     getColumnValue(data["receipt_link"]),
		"remarks":          getColumnValue(data["remarks"]),
		"requested_date":   getColumnValue(data["requested_date"]),
		"requestor_id":     getColumnValue(data["requestor_id"]),
		"requestor_name":   getColumnValue(data["requestor_name"]),
		"requestor_role":   getColumnValue(data["requestor_role"]),
		"requestor_wallet_id": getColumnValue(data["requestor_wallet_id"]),
		"status":           getColumnValue(data["status"]),
		"status_desc":      getColumnValue(data["status_desc"]),
		"sub_requestor_id": getColumnValue(data["sub_requestor_id"]),
		"sub_requestor_name": getColumnValue(data["sub_requestor_name"]),
		"sub_requestor_role": getColumnValue(data["sub_requestor_role"]),
		"sub_status":       getColumnValue(data["sub_status"]),
		"sync_time":        getColumnValue(data["sync_time"]),
		"terminal_description": getColumnValue(data["terminal_description"]),
		"terminal_id":      getColumnValue(data["terminal_id"]),
		"transaction_id":   getColumnValue(data["transaction_id"]),
		"transaction_mode": getColumnValue(data["transaction_mode"]),
		"transaction_operation": getColumnValue(data["transaction_operation"]),
		"transaction_type": getColumnValue(data["transaction_type"]),
		"updated_date":     getColumnValue(data["updated_date"]),
		"utr_number":       getColumnValue(data["utr_number"]),
		"wallet_id":        getColumnValue(data["wallet_id"]),
		"wallet_ledger_id": getColumnValue(data["wallet_ledger_id"]),
		"wallet_reversed":  getColumnValue(data["wallet_reversed"]),
		"wallet_status":    getColumnValue(data["wallet_status"]),
		"wallet_status_desc": getColumnValue(data["wallet_status_desc"]),
		"parent_transaction_id": getColumnValue(data["parent_transaction_id"]),
		"reversal_type":    getColumnValue(data["reversal_type"]),
		"refund_reason":    getColumnValue(data["refund_reason"]),
		"card_number":      getColumnValue(data["card_number"]),
		"settlement_date":  getColumnValue(data["settlement_date"]),
		"settlement_time":  getColumnValue(data["settlement_time"]),
	}

	fmt.Println("Filtered Data:", filteredData)

	return filteredData, nil
}



func FilterDataForUserCreation(data map[string]interface{}, bookMark string) (map[string]interface{}, error) {
	nullArr := map[interface{}]bool{nil: true, "null": true, "" : true}

	getColumnValue := func(column interface{}) interface{} {
		if _, exists := nullArr[column]; exists {
			return nil
		}
		return column
	}

	filteredData := map[string]interface{}{
		"user_id":                     getColumnValue(data["user_id"]),
		"address":                     getColumnValue(data["address"]),
		"admin_id":                    getColumnValue(data["admin_id"]),
		"admin_name":                  getColumnValue(data["admin_name"]),
		"admin_role":                  getColumnValue(data["admin_role"]),
		"approved_by":                 getColumnValue(data["approved_by"]),
		"approver_id":                 getColumnValue(data["approver_id"]),
		"approver_role":               getColumnValue(data["approver_role"]),
		"city":                        getColumnValue(data["city"]),
		"country":                     getColumnValue(data["country"]),
		"created_date":                getColumnValue(data["created_date"]),
		"creator_id":                  getColumnValue(data["creator_id"]),
		"creator_role":                getColumnValue(data["creator_role"]),
		"date_of_birth":               getColumnValue(data["date_of_birth"]),
		"email":                       getColumnValue(data["email"]),
		"first_name":                  getColumnValue(data["first_name"]),
		"gender":                      getColumnValue(data["gender"]),
		"is_active":                   getColumnValue(data["is_active"]),
		"is_block":                    getColumnValue(data["is_block"]),
		"kyc_number":                  getColumnValue(data["kyc_number"]),
		"kyc_status":                  getColumnValue(data["kyc_status"]),
		"kyc_type":                    getColumnValue(data["kyc_type"]),
		"last_name":                   getColumnValue(data["last_name"]),
		"latlong":                     getColumnValue(data["latlong"]),
		"mobile_number":               getColumnValue(data["mobile_number"]),
		"pincode":                     getColumnValue(data["pincode"]),
		"self_onboard":                getColumnValue(data["self_onboard"]),
		"state":                       getColumnValue(data["state"]),
		"updated_date":                getColumnValue(data["updated_date"]),
		"user_name":                   getColumnValue(data["user_name"]),
		"user_role":                   getColumnValue(data["user_role"]),
		"wallet_id":                   getColumnValue(data["wallet_id"]),
		"wallet_status":               getColumnValue(data["wallet_status"]),
	}

	fmt.Println("Filtered User Data:", filteredData)
	return filteredData,nil
}
