package main

import (
	"fmt"
	"unsafe"

	"github.com/go-pg/pg"
)

type ActionUser struct {
	tableName struct{} `pg:"customer.leads_action_user,discard_unknown_columns" mapstructure:",omitempty"`
	ID        string   `json:"-" mapstructure:",omitempty"`
	Name      string   `json:"name" mapstructure:"name"`
	Email     string   `json:"email" mapstructure:"email"`
}

type Banks struct {
	tableName               struct{}    `pg:"customer.leads_bank,discard_unknown_columns" mapstructure:",omitempty"`
	IsActive                bool        `json:"is_active" mapstructure:"is_active"`
	IsLeadAccountHolder     bool        `json:"is_lead_account_holder" mapstructure:"is_lead_account_holder"`
	IsMainAccount           bool        `json:"is_main_account" mapstructure:"is_main_account"`
	LeadsID                 string      `json:"-" mapstructure:",omitempty"`
	ID                      string      `json:"-" pg:"id" mapstructure:",omitempty"`
	Name                    string      `json:"name" mapstructure:"name"`
	AccountHolderName       string      `json:"account_holder_name" mapstructure:"account_holder_name"`
	AccountRelationship     string      `json:"account_relationship" mapstructure:"account_relationship"`
	AccountNumber           string      `json:"account_number" mapstructure:"account_number"`
	BankIrisCode            string      `json:"bank_iris_code" mapstructure:"bank_iris_code"`
	AccountProofOfOwnership string      `json:"account_proof_of_ownership" mapstructure:"account_proof_of_ownership"`
	ValidateStatus          string      `json:"validate_status" mapstructure:"validate_status"`
	ValidateBy              string      `json:"validate_by" mapstructure:"validate_by"`
	CreatedByApp            string      `json:"-" pg:"created_by_app" mapstructure:",omitempty"`
	CreatedByID             string      `json:"-" pg:"created_by_id" mapstructure:",omitempty"`
	UpdatedByID             string      `json:"-" pg:"updated_by_id" mapstructure:",omitempty"`
	UpdatedByApp            string      `json:"-" pg:"updated_by_app" mapstructure:",omitempty"`
	CreatedBy               ActionUser  `json:"createdBy" pg:"rel:has-one,fk:created_by_id" mapstructure:"createdBy"`
	UpdatedBy               ActionUser  `json:"updatedBy" pg:"rel:has-one,fk:updated_by_id" mapstructure:"updatedBy"`
	CreatedAt               pg.NullTime `json:"createdAt" mapstructure:"createdAt"`
	ValidateAt              pg.NullTime `json:"validate_at" pg:"validate_at" mapstructure:"validate_at"`
	UpdatedAt               pg.NullTime `json:"updatedAt" mapstructure:"updatedAt"`
}

type Ponds struct {
	tableName           struct{}    `pg:"customer.leads_locations_ponds,discard_unknown_columns" mapstructure:",omitempty"`
	Latitude            float64     `json:"latitude" mapstructure:"latitude"`
	Longitude           float64     `json:"longitude" mapstructure:"longitude"`
	Large               float64     `json:"large" mapstructure:"large"`
	Depth               float64     `json:"depth" mapstructure:"depth"`
	IsTemporaryDisabled bool        `json:"is_temporary_disabled" mapstructure:"is_temporary_disabled"`
	PondCapacity        *int        `json:"pond_capacity" mapstructure:"pond_capacity"`
	IsActive            *bool       `json:"is_active" mapstructure:"is_active"`
	IsMpi               *bool       `json:"is_mpi" mapstructure:"is_mpi"`
	LocationsID         string      `json:"locations_id" mapstructure:"locations_id"`
	UUID                string      `json:"uuid" mapstructure:"uuid"`
	Name                string      `json:"name" binding:"required" mapstructure:"name"`
	PondImage           string      `json:"pond_image" mapstructure:"pond_image"`
	PondType            string      `json:"pond_type" mapstructure:"pond_type"`
	PondShape           string      `json:"pond_shape" mapstructure:"pond_shape"`
	LinkMap             string      `json:"link_map" mapstructure:"link_map"`
	CreatedByID         string      `json:"-" pg:"created_by_id" mapstructure:"created_by"`
	CreatedByApp        string      `json:"-" pg:"created_by_app" mapstructure:"created_by_app"`
	UpdatedByID         string      `json:"-" pg:"updated_by_id" mapstructure:"updated_by"`
	UpdatedByApp        string      `json:"-" pg:"updated_by_app" mapstructure:"updated_by_app"`
	DisabledUntil       pg.NullTime `json:"disabled_until" mapstructure:"disabled_until"`
	CreatedAt           pg.NullTime `json:"createdAt" mapstructure:"created_at"`
	UpdatedAt           pg.NullTime `json:"updatedAt" mapstructure:"updated_at"`
	CreatedBy           ActionUser  `json:"createdBy" pg:"rel:has-one,fk:created_by_id" mapstructure:"createdBy"`
	UpdatedBy           ActionUser  `json:"updatedBy" pg:"rel:has-one,fk:updated_by_id" mapstructure:"updatedBy"`
}

func main() {
	b := Banks{}
	fmt.Println(unsafe.Sizeof(b)) //16

	fmt.Println("-----------------------")

	p := Ponds{}
	fmt.Println(unsafe.Sizeof(p)) //16

}
