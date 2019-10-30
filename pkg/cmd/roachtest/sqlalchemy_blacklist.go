// Copyright 2019 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package main

var sqlAlchemyBlacklists = blacklistsForVersion{
	{"v2.1", "sqlAlchemyBlacklist", sqlAlchemyBlacklist, "sqlAlchemyIgnoreList", sqlAlchemyIgnoreList},
	{"v19.1", "sqlAlchemyBlacklist", sqlAlchemyBlacklist, "sqlAlchemyIgnoreList", sqlAlchemyIgnoreList},
	{"v19.2", "sqlAlchemyBlacklist", sqlAlchemyBlacklist, "sqlAlchemyIgnoreList", sqlAlchemyIgnoreList},
}

var sqlAlchemyBlacklist = blacklist{
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_autoincrement_col":                  "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_deprecated_get_primary_keys":        "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_dialect_initialize":                 "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_get_check_constraints":              "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_get_check_constraints_schema":       "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_get_columns":                        "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_get_columns_with_schema":            "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_get_comments":                       "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_get_comments_with_schema":           "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_get_default_schema_name":            "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_get_foreign_key_options_ondelete":   "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_get_foreign_key_options_onupdate":   "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_get_foreign_keys":                   "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_get_foreign_keys_with_schema":       "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_get_indexes":                        "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_get_indexes_with_schema":            "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_get_inter_schema_foreign_keys":      "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_get_noncol_index_no_pk":             "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_get_noncol_index_pk":                "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_get_pk_constraint":                  "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_get_pk_constraint_with_schema":      "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_get_schema_names":                   "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_get_table_names":                    "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_get_table_names_fks":                "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_get_table_names_with_schema":        "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_get_table_oid":                      "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_get_table_oid_with_schema":          "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_get_tables_and_views":               "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_get_temp_table_columns":             "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_get_temp_table_indexes":             "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_get_temp_table_names":               "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_get_temp_table_unique_constraints":  "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_get_temp_view_columns":              "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_get_temp_view_names":                "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_get_unique_constraints":             "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_get_unique_constraints_with_schema": "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_get_view_columns":                   "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_get_view_columns_with_schema":       "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_get_view_definition":                "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_get_view_definition_with_schema":    "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_get_view_names":                     "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_get_view_names_with_schema":         "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_nullable_reflection":                "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_numeric_reflection":                 "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_reflect_expression_based_indexes":   "5807",
	"test/dialect/test_suite.py::ComponentReflectionTest_cockroachdb+psycopg2_9_5_0::test_varchar_reflection":                 "5807",
	"test/dialect/test_suite.py::DateTimeCoercedToDateTimeTest_cockroachdb+psycopg2_9_5_0::test_null_bound_comparison":        "36179",
	"test/dialect/test_suite.py::DateTimeCoercedToDateTimeTest_cockroachdb+psycopg2_9_5_0::test_round_trip":                   "36179",
	"test/dialect/test_suite.py::ExpandingBoundInTest_cockroachdb+psycopg2_9_5_0::test_null_in_empty_set_is_false":            "41596",
	"test/dialect/test_suite.py::HasTableTest_cockroachdb+psycopg2_9_5_0::test_has_table":                                     "26443",
	"test/dialect/test_suite.py::HasTableTest_cockroachdb+psycopg2_9_5_0::test_has_table_schema":                              "26443",
	"test/dialect/test_suite.py::LastrowidTest_cockroachdb+psycopg2_9_5_0::test_autoincrement_on_insert":                      "41690",
	"test/dialect/test_suite.py::ReturningTest_cockroachdb+psycopg2_9_5_0::test_autoincrement_on_insert_implicit_returning":   "41690",
	"test/dialect/test_suite.py::ServerSideCursorsTest_cockroachdb+psycopg2_9_5_0::test_aliases_and_ss":                       "41412",
	"test/dialect/test_suite.py::ServerSideCursorsTest_cockroachdb+psycopg2_9_5_0::test_conn_option":                          "41412",
	"test/dialect/test_suite.py::ServerSideCursorsTest_cockroachdb+psycopg2_9_5_0::test_for_update_expr":                      "41412",
	"test/dialect/test_suite.py::ServerSideCursorsTest_cockroachdb+psycopg2_9_5_0::test_for_update_string":                    "41412",
	"test/dialect/test_suite.py::ServerSideCursorsTest_cockroachdb+psycopg2_9_5_0::test_global_expr":                          "41412",
	"test/dialect/test_suite.py::ServerSideCursorsTest_cockroachdb+psycopg2_9_5_0::test_global_string":                        "41412",
	"test/dialect/test_suite.py::ServerSideCursorsTest_cockroachdb+psycopg2_9_5_0::test_global_text":                          "41412",
	"test/dialect/test_suite.py::ServerSideCursorsTest_cockroachdb+psycopg2_9_5_0::test_roundtrip":                            "41412",
	"test/dialect/test_suite.py::ServerSideCursorsTest_cockroachdb+psycopg2_9_5_0::test_stmt_option":                          "41412",
	"test/dialect/test_suite.py::ServerSideCursorsTest_cockroachdb+psycopg2_9_5_0::test_stmt_option_disabled":                 "41412",
	"test/dialect/test_suite.py::ServerSideCursorsTest_cockroachdb+psycopg2_9_5_0::test_text_ss_option":                       "41412",
	"test/dialect/test_suite.py::TableDDLTest_cockroachdb+psycopg2_9_5_0::test_create_table_schema":                           "26443",
}

var sqlAlchemyIgnoreList = blacklist{
	"test/dialect/test_suite.py::TableDDLTest_cockroachdb+psycopg2_9_5_0::test_create_table": "flaky",
}
