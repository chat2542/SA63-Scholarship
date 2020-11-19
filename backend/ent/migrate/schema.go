// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// ScholarshipRequestsColumns holds the columns for the "scholarship_requests" table.
	ScholarshipRequestsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "scholarshipinformation_scholarship_request", Type: field.TypeInt, Nullable: true},
		{Name: "scholarshiptype_scholarship_request", Type: field.TypeInt, Nullable: true},
		{Name: "semester_scholarship_request", Type: field.TypeInt, Nullable: true},
		{Name: "user_scholarship_request", Type: field.TypeInt, Nullable: true},
	}
	// ScholarshipRequestsTable holds the schema information for the "scholarship_requests" table.
	ScholarshipRequestsTable = &schema.Table{
		Name:       "scholarship_requests",
		Columns:    ScholarshipRequestsColumns,
		PrimaryKey: []*schema.Column{ScholarshipRequestsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "scholarship_requests_scholarshipinformations_ScholarshipRequest",
				Columns: []*schema.Column{ScholarshipRequestsColumns[1]},

				RefColumns: []*schema.Column{ScholarshipinformationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "scholarship_requests_scholarshiptypes_ScholarshipRequest",
				Columns: []*schema.Column{ScholarshipRequestsColumns[2]},

				RefColumns: []*schema.Column{ScholarshiptypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "scholarship_requests_semesters_ScholarshipRequest",
				Columns: []*schema.Column{ScholarshipRequestsColumns[3]},

				RefColumns: []*schema.Column{SemestersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "scholarship_requests_users_ScholarshipRequest",
				Columns: []*schema.Column{ScholarshipRequestsColumns[4]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ScholarshipinformationsColumns holds the columns for the "scholarshipinformations" table.
	ScholarshipinformationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "scholarship_name", Type: field.TypeString, Unique: true},
		{Name: "scholarshiptype_scholarshipinformation", Type: field.TypeInt, Nullable: true},
	}
	// ScholarshipinformationsTable holds the schema information for the "scholarshipinformations" table.
	ScholarshipinformationsTable = &schema.Table{
		Name:       "scholarshipinformations",
		Columns:    ScholarshipinformationsColumns,
		PrimaryKey: []*schema.Column{ScholarshipinformationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "scholarshipinformations_scholarshiptypes_Scholarshipinformation",
				Columns: []*schema.Column{ScholarshipinformationsColumns[2]},

				RefColumns: []*schema.Column{ScholarshiptypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ScholarshiptypesColumns holds the columns for the "scholarshiptypes" table.
	ScholarshiptypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type_name", Type: field.TypeString, Unique: true},
	}
	// ScholarshiptypesTable holds the schema information for the "scholarshiptypes" table.
	ScholarshiptypesTable = &schema.Table{
		Name:        "scholarshiptypes",
		Columns:     ScholarshiptypesColumns,
		PrimaryKey:  []*schema.Column{ScholarshiptypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// SemestersColumns holds the columns for the "semesters" table.
	SemestersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "semester", Type: field.TypeString, Unique: true},
	}
	// SemestersTable holds the schema information for the "semesters" table.
	SemestersTable = &schema.Table{
		Name:        "semesters",
		Columns:     SemestersColumns,
		PrimaryKey:  []*schema.Column{SemestersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "email", Type: field.TypeString, Unique: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ScholarshipRequestsTable,
		ScholarshipinformationsTable,
		ScholarshiptypesTable,
		SemestersTable,
		UsersTable,
	}
)

func init() {
	ScholarshipRequestsTable.ForeignKeys[0].RefTable = ScholarshipinformationsTable
	ScholarshipRequestsTable.ForeignKeys[1].RefTable = ScholarshiptypesTable
	ScholarshipRequestsTable.ForeignKeys[2].RefTable = SemestersTable
	ScholarshipRequestsTable.ForeignKeys[3].RefTable = UsersTable
	ScholarshipinformationsTable.ForeignKeys[0].RefTable = ScholarshiptypesTable
}
