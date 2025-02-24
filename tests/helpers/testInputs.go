package helpers

var BatchPayload = `
{
    "batch": [
        {
            "userId": "identified user id",
            "anonymousId": "anon-id-new",
            "event": "Product Purchased new",
            "type": "track",
            "properties": {
                "name": "Shirt",
                "revenue": 4.99
            },
            "context": {
                "ip": "14.5.67.21",
                "library": {
                    "name": "http"
                }
            },
            "timestamp": "2020-02-02T00:23:09.544Z"
        }
    ]
}
`
var LoadedJson = `
{
  "anonymous_id": "af0d6ca8-ac18-42dc-ae98-cc40ba5b7c4d",
  "category": "Demo Category",
  "context_app_build": "1",
  "context_app_name": "RudderAndroidClient",
  "context_app_namespace": "com.rudderlabs.android.sdk",
  "context_app_version": "1.0",
  "context_device_id": "49e4bdd1c280bc00",
  "context_device_manufacturer": "Google",
  "context_device_model": "Android SDK built for x86",
  "context_device_name": "generic_x86",
  "context_ip": "[::1]:55682",
  "context_locale": "en-US",
  "context_network_carrier": "Android",
  "context_screen_density": 420,
  "context_screen_height": 1794,
  "context_screen_width": 1080,
  "context_traits_anonymousId": "49e4bdd1c280bc00",
  "context_user_agent": "Dalvik/2.1.0 (Linux; U; Android 9; Android SDK built for x86 Build/PSR1.180720.075)",
  "event": "ginkgo",
  "event_text": "ginkgo",
  "id": "74b582ef-7663-4cb3-8666-2273fe5a8ccb",
  "label": "Demo Label",
  "original_timestamp": "2019-08-12T05:08:30.909Z",
  "received_at": "2020-03-25T15:51:24.971Z",
  "sent_at": "2019-08-12T05:08:30.909Z",
  "timestamp": "2020-03-25T15:51:24.971Z",
  "uuid_ts": "2020-03-25 21:32:29 Z",
  "value": 5
}
`

var IdentifyPayload = `
{
	"anonymousId": "49e4bdd1c280bc00",
	"messageId": "msgasdfadsf2er34adfsdf2",
	"type": "identify",
	"userId": "98234023840234",
	"traits": {
	  "name": "Chandra",
	  "email": "chandra@rudderlabs.com",
	  "org": "rudder"
	}
}
`

var AliasPayload = `
{
	"anonymousId": "49e4bdd1c280bc00",
	"messageId": "msgasdfadsf2er34adfsdf3",
	"type": "alias",
	"previousId": "chandra@rudderlabs.com",
	"userId": "98234023840234"
}
`

var TrackPayload = `
{
	"anonymousId": "49e4bdd1c280bc00",
	"messageId": "msgasdfadsf2er34adfsdf4",
	"type": "track",
	"event": "test event",
	"properties": {
		"name": "Chandra"
	}
}
`

var GroupPayload = `
{
	"anonymousId": "49e4bdd1c280bc00",
	"messageId": "msgasdfadsf2er34adfsdf5",
	"type": "group",
	"groupId": "98234023840234adf2e232",
	"traits": {
		"name": "Chandra",
		"email": "chandra@rudderlabs.com",
		"org": "rudder"
	}
}
`

var PagePayload = `
{
	"anonymousId": "49e4bdd1c280bc00",
	"messageId": "msgasdfadsf2er34adfsdf6",
	"type": "page",
	"name": "Hello",
	"properties": {
	  "title": "Welcome to rudder",
	  "url": "http://www.rudderstack.com"
	}
  }
`

var ScreenPayload = `
{
	"anonymousId": "49e4bdd1c280bc00",
	"messageId": "msgasdfadsf2er34adfsdf7",
	"type": "screen",
	"name": "Hello",
	"properties": {
		"title": "Welcome to rudder"
	}
}
`

var WarehouseBatchPayload = `
{
	"batch": [
		{
		"anonymousId": "49e4bdd1c280bc00",
		"messageId": "msgasdfadsf2er34adfsdf1",
		"channel": "android-sdk",
		"destination_props": {
			"AF": {
			"af_uid": "1566363489499-3377330514807116178"
			}
		},
		"context": {
			"app": {
			"build": "1",
			"name": "RudderAndroidClient",
			"namespace": "com.rudderlabs.android.sdk",
			"version": "1.0"
			},
			"device": {
			"id": "49e4bdd1c280bc00",
			"manufacturer": "Google",
			"model": "Android SDK built for x86",
			"name": "generic_x86"
			},
			"locale": "en-US",
			"network": {
			"carrier": "Android"
			},
			"screen": {
			"density": 420,
			"height": 1794,
			"width": 1080
			},
			"traits": {
			"anonymousId": "49e4bdd1c280bc00"
			},
			"user_agent": "Dalvik/2.1.0 (Linux; U; An  droid 9; Android SDK built for x86 Build/PSR1.180720.075)"
		},
		"event": "Demo Track",
		"integrations": {
			"All": true
		},
		"properties": {
			"label": "Demo Label",
			"category": "Demo Category",
			"value": 5,
			"property1":"test",
			"property2":"test",
			"property3":"test",
			"property4":"test",
			"property5":"test"
		},
		"type": "track",
		"originalTimestamp": "2019-08-12T05:08:30.909Z",
		"sentAt": "2019-08-12T05:08:30.909Z"
		}
	]
}
`

var POSTGRESSchema = map[string]map[string]string{
	"ginkgo": {
		"anonymous_id":                "string",
		"category":                    "string",
		"channel":                     "string",
		"context_app_build":           "string",
		"context_app_name":            "string",
		"context_app_namespace":       "string",
		"context_app_version":         "string",
		"context_device_id":           "string",
		"context_device_manufacturer": "string",
		"context_device_model":        "string",
		"context_device_name":         "string",
		"context_ip":                  "string",
		"context_locale":              "string",
		"context_network_carrier":     "string",
		"context_request_ip":          "string",
		"context_screen_density":      "int",
		"context_screen_height":       "int",
		"context_screen_width":        "int",
		"context_traits_anonymous_id": "string",
		"context_user_agent":          "string",
		"event":                       "string",
		"event_text":                  "string",
		"id":                          "string",
		"label":                       "string",
		"name":                        "string",
		"original_timestamp":          "datetime",
		"property_1":                  "string",
		"property_2":                  "string",
		"property_3":                  "string",
		"property_4":                  "string",
		"property_5":                  "string",
		"received_at":                 "datetime",
		"sent_at":                     "datetime",
		"timestamp":                   "datetime",
		"uuid_ts":                     "datetime",
		"value":                       "int",
	},
	"tracks": {
		"anonymous_id":                "string",
		"channel":                     "string",
		"context_app_build":           "string",
		"context_app_name":            "string",
		"context_app_namespace":       "string",
		"context_app_version":         "string",
		"context_device_id":           "string",
		"context_device_manufacturer": "string",
		"context_device_model":        "string",
		"context_device_name":         "string",
		"context_ip":                  "string",
		"context_locale":              "string",
		"context_network_carrier":     "string",
		"context_request_ip":          "string",
		"context_screen_density":      "int",
		"context_screen_height":       "int",
		"context_screen_width":        "int",
		"context_traits_anonymous_id": "string",
		"context_user_agent":          "string",
		"event":                       "string",
		"event_text":                  "string",
		"id":                          "string",
		"original_timestamp":          "datetime",
		"received_at":                 "datetime",
		"sent_at":                     "datetime",
		"timestamp":                   "datetime",
		"uuid_ts":                     "datetime",
	},
}
var RedshiftSchema = map[string]map[string]string{
	"ginkgo": {
		"_from":                       "string",
		"_join":                       "string",
		"_order":                      "string",
		"_select":                     "string",
		"_where":                      "string",
		"anonymous_id":                "string",
		"category":                    "string",
		"channel":                     "string",
		"context_app_build":           "string",
		"context_app_name":            "string",
		"context_app_namespace":       "string",
		"context_app_version":         "string",
		"context_device_id":           "string",
		"context_device_manufacturer": "string",
		"context_device_model":        "string",
		"context_device_name":         "string",
		"context_ip":                  "string",
		"context_locale":              "string",
		"context_network_carrier":     "string",
		"context_request_ip":          "string",
		"context_screen_density":      "int",
		"context_screen_height":       "int",
		"context_screen_width":        "int",
		"context_traits_anonymous_id": "string",
		"context_user_agent":          "string",
		"event":                       "string",
		"event_text":                  "string",
		"id":                          "string",
		"label":                       "string",
		"name":                        "string",
		"original_timestamp":          "datetime",
		"property_1":                  "string",
		"property_2":                  "string",
		"property_3":                  "string",
		"property_4":                  "string",
		"property_5":                  "string",
		"received_at":                 "datetime",
		"sent_at":                     "datetime",
		"timestamp":                   "datetime",
		"uuid_ts":                     "datetime",
		"value":                       "int",
	},
	"tracks": {
		"anonymous_id":                "string",
		"channel":                     "string",
		"context_app_build":           "string",
		"context_app_name":            "string",
		"context_app_namespace":       "string",
		"context_app_version":         "string",
		"context_device_id":           "string",
		"context_device_manufacturer": "string",
		"context_device_model":        "string",
		"context_device_name":         "string",
		"context_ip":                  "string",
		"context_library_name":        "string",
		"context_locale":              "string",
		"context_network_carrier":     "string",
		"context_passed_ip":           "string",
		"context_request_ip":          "string",
		"context_screen_density":      "int",
		"context_screen_height":       "int",
		"context_screen_width":        "int",
		"context_traits_anonymous_id": "string",
		"context_user_agent":          "string",
		"event":                       "string",
		"event_text":                  "string",
		"id":                          "string",
		"original_timestamp":          "datetime",
		"received_at":                 "datetime",
		"sent_at":                     "datetime",
		"timestamp":                   "datetime",
		"user_id":                     "string",
		"uuid_ts":                     "datetime",
	},
}
var ReservedKeywordsRedshiftSchema = map[string]map[string]string{
	"ginkgo": {
		"_from":                       "string",
		"_join":                       "string",
		"_order":                      "string",
		"_select":                     "string",
		"_where":                      "string",
		"anonymous_id":                "string",
		"category":                    "string",
		"channel":                     "string",
		"context_app_build":           "string",
		"context_app_name":            "string",
		"context_app_namespace":       "string",
		"context_app_version":         "string",
		"context_device_id":           "string",
		"context_device_manufacturer": "string",
		"context_device_model":        "string",
		"context_device_name":         "string",
		"context_ip":                  "string",
		"context_locale":              "string",
		"context_network_carrier":     "string",
		"context_request_ip":          "string",
		"context_screen_density":      "int",
		"context_screen_height":       "int",
		"context_screen_width":        "int",
		"context_traits_anonymous_id": "string",
		"context_user_agent":          "string",
		"event":                       "string",
		"event_text":                  "string",
		"id":                          "string",
		"label":                       "string",
		"name":                        "string",
		"original_timestamp":          "datetime",
		"property_1":                  "string",
		"property_2":                  "string",
		"property_3":                  "string",
		"property_4":                  "string",
		"property_5":                  "string",
		"received_at":                 "datetime",
		"sent_at":                     "datetime",
		"timestamp":                   "datetime",
		"uuid_ts":                     "datetime",
		"value":                       "int",
	},
	"tracks": {
		"anonymous_id":                "string",
		"channel":                     "string",
		"context_app_build":           "string",
		"context_app_name":            "string",
		"context_app_namespace":       "string",
		"context_app_version":         "string",
		"context_device_id":           "string",
		"context_device_manufacturer": "string",
		"context_device_model":        "string",
		"context_device_name":         "string",
		"context_ip":                  "string",
		"context_library_name":        "string",
		"context_locale":              "string",
		"context_network_carrier":     "string",
		"context_passed_ip":           "string",
		"context_request_ip":          "string",
		"context_screen_density":      "int",
		"context_screen_height":       "int",
		"context_screen_width":        "int",
		"context_traits_anonymous_id": "string",
		"context_user_agent":          "string",
		"event":                       "string",
		"event_text":                  "string",
		"id":                          "string",
		"original_timestamp":          "datetime",
		"received_at":                 "datetime",
		"sent_at":                     "datetime",
		"timestamp":                   "datetime",
		"user_id":                     "string",
		"uuid_ts":                     "datetime",
	},
}

var BigQuerySchema = map[string]map[string]string{
	"ginkgo": {
		"_from":                       "string",
		"_join":                       "string",
		"_order":                      "string",
		"_select":                     "string",
		"_where":                      "string",
		"anonymous_id":                "string",
		"category":                    "string",
		"channel":                     "string",
		"context_app_build":           "string",
		"context_app_name":            "string",
		"context_app_namespace":       "string",
		"context_app_version":         "string",
		"context_device_id":           "string",
		"context_device_manufacturer": "string",
		"context_device_model":        "string",
		"context_device_name":         "string",
		"context_ip":                  "string",
		"context_locale":              "string",
		"context_network_carrier":     "string",
		"context_request_ip":          "string",
		"context_screen_density":      "int",
		"context_screen_height":       "int",
		"context_screen_width":        "int",
		"context_traits_anonymous_id": "string",
		"context_user_agent":          "string",
		"event":                       "string",
		"event_text":                  "string",
		"id":                          "string",
		"label":                       "string",
		"loaded_at":                   "datetime",
		"name":                        "string",
		"original_timestamp":          "datetime",
		"property_1":                  "string",
		"property_2":                  "string",
		"property_3":                  "string",
		"property_4":                  "string",
		"property_5":                  "string",
		"received_at":                 "datetime",
		"sent_at":                     "datetime",
		"timestamp":                   "datetime",
		"uuid_ts":                     "datetime",
		"value":                       "int",
	},
	"tracks": {
		"anonymous_id":                "string",
		"channel":                     "string",
		"context_app_build":           "string",
		"context_app_name":            "string",
		"context_app_namespace":       "string",
		"context_app_version":         "string",
		"context_device_id":           "string",
		"context_device_manufacturer": "string",
		"context_device_model":        "string",
		"context_device_name":         "string",
		"context_ip":                  "string",
		"context_locale":              "string",
		"context_network_carrier":     "string",
		"context_request_ip":          "string",
		"context_screen_density":      "int",
		"context_screen_height":       "int",
		"context_screen_width":        "int",
		"context_traits_anonymous_id": "string",
		"context_user_agent":          "string",
		"event":                       "string",
		"event_text":                  "string",
		"id":                          "string",
		"loaded_at":                   "datetime",
		"original_timestamp":          "datetime",
		"received_at":                 "datetime",
		"sent_at":                     "datetime",
		"timestamp":                   "datetime",
		"uuid_ts":                     "datetime",
	},
}
var ReserverKeyWordsBigQuerySchema = map[string]map[string]string{
	"ginkgo": {
		"_from":                       "string",
		"_join":                       "string",
		"_order":                      "string",
		"_select":                     "string",
		"_where":                      "string",
		"anonymous_id":                "string",
		"category":                    "string",
		"channel":                     "string",
		"context_app_build":           "string",
		"context_app_name":            "string",
		"context_app_namespace":       "string",
		"context_app_version":         "string",
		"context_device_id":           "string",
		"context_device_manufacturer": "string",
		"context_device_model":        "string",
		"context_device_name":         "string",
		"context_ip":                  "string",
		"context_locale":              "string",
		"context_network_carrier":     "string",
		"context_request_ip":          "string",
		"context_screen_density":      "int",
		"context_screen_height":       "int",
		"context_screen_width":        "int",
		"context_traits_anonymous_id": "string",
		"context_user_agent":          "string",
		"event":                       "string",
		"event_text":                  "string",
		"id":                          "string",
		"label":                       "string",
		"loaded_at":                   "datetime",
		"name":                        "string",
		"original_timestamp":          "datetime",
		"property_1":                  "string",
		"property_2":                  "string",
		"property_3":                  "string",
		"property_4":                  "string",
		"property_5":                  "string",
		"received_at":                 "datetime",
		"sent_at":                     "datetime",
		"timestamp":                   "datetime",
		"uuid_ts":                     "datetime",
		"value":                       "int",
	},
	"tracks": {
		"anonymous_id":                "string",
		"channel":                     "string",
		"context_app_build":           "string",
		"context_app_name":            "string",
		"context_app_namespace":       "string",
		"context_app_version":         "string",
		"context_device_id":           "string",
		"context_device_manufacturer": "string",
		"context_device_model":        "string",
		"context_device_name":         "string",
		"context_ip":                  "string",
		"context_locale":              "string",
		"context_network_carrier":     "string",
		"context_request_ip":          "string",
		"context_screen_density":      "int",
		"context_screen_height":       "int",
		"context_screen_width":        "int",
		"context_traits_anonymous_id": "string",
		"context_user_agent":          "string",
		"event":                       "string",
		"event_text":                  "string",
		"id":                          "string",
		"loaded_at":                   "datetime",
		"original_timestamp":          "datetime",
		"received_at":                 "datetime",
		"sent_at":                     "datetime",
		"timestamp":                   "datetime",
		"uuid_ts":                     "datetime",
	},
}

var SnowflakeSchema = map[string]map[string]string{
	"GINKGO": {
		"_FROM":                       "string",
		"_JOIN":                       "string",
		"_ORDER":                      "string",
		"_SELECT":                     "string",
		"_WHERE":                      "string",
		"ANONYMOUS_ID":                "string",
		"CATEGORY":                    "string",
		"CHANNEL":                     "string",
		"CONTEXT_APP_BUILD":           "string",
		"CONTEXT_APP_NAME":            "string",
		"CONTEXT_APP_NAMESPACE":       "string",
		"CONTEXT_APP_VERSION":         "string",
		"CONTEXT_DEVICE_ID":           "string",
		"CONTEXT_DEVICE_MANUFACTURER": "string",
		"CONTEXT_DEVICE_MODEL":        "string",
		"CONTEXT_DEVICE_NAME":         "string",
		"CONTEXT_IP":                  "string",
		"CONTEXT_LOCALE":              "string",
		"CONTEXT_NETWORK_CARRIER":     "string",
		"CONTEXT_REQUEST_IP":          "string",
		"CONTEXT_SCREEN_DENSITY":      "int",
		"CONTEXT_SCREEN_HEIGHT":       "int",
		"CONTEXT_SCREEN_WIDTH":        "int",
		"CONTEXT_TRAITS_ANONYMOUS_ID": "string",
		"CONTEXT_USER_AGENT":          "string",
		"EVENT":                       "string",
		"EVENT_TEXT":                  "string",
		"ID":                          "string",
		"LABEL":                       "string",
		"NAME":                        "string",
		"ORIGINAL_TIMESTAMP":          "datetime",
		"PROPERTY_1":                  "string",
		"PROPERTY_2":                  "string",
		"PROPERTY_3":                  "string",
		"PROPERTY_4":                  "string",
		"PROPERTY_5":                  "string",
		"RECEIVED_AT":                 "datetime",
		"SENT_AT":                     "datetime",
		"TIMESTAMP":                   "datetime",
		"UUID_TS":                     "datetime",
		"VALUE":                       "int",
	},
	"TRACKS": {
		"ANONYMOUS_ID":                "string",
		"CHANNEL":                     "string",
		"CONTEXT_APP_BUILD":           "string",
		"CONTEXT_APP_NAME":            "string",
		"CONTEXT_APP_NAMESPACE":       "string",
		"CONTEXT_APP_VERSION":         "string",
		"CONTEXT_DEVICE_ID":           "string",
		"CONTEXT_DEVICE_MANUFACTURER": "string",
		"CONTEXT_DEVICE_MODEL":        "string",
		"CONTEXT_DEVICE_NAME":         "string",
		"CONTEXT_IP":                  "string",
		"CONTEXT_LOCALE":              "string",
		"CONTEXT_NETWORK_CARRIER":     "string",
		"CONTEXT_REQUEST_IP":          "string",
		"CONTEXT_SCREEN_DENSITY":      "int",
		"CONTEXT_SCREEN_HEIGHT":       "int",
		"CONTEXT_SCREEN_WIDTH":        "int",
		"CONTEXT_TRAITS_ANONYMOUS_ID": "string",
		"CONTEXT_USER_AGENT":          "string",
		"EVENT":                       "string",
		"EVENT_TEXT":                  "string",
		"ID":                          "string",
		"ORIGINAL_TIMESTAMP":          "datetime",
		"RECEIVED_AT":                 "datetime",
		"SENT_AT":                     "datetime",
		"TIMESTAMP":                   "datetime",
		"UUID_TS":                     "datetime",
	},
}
var ReservedKeywordsSnowflakeSchema = map[string]map[string]string{
	"GINKGO": {
		"_FROM":                       "string",
		"_JOIN":                       "string",
		"_ORDER":                      "string",
		"_SELECT":                     "string",
		"_WHERE":                      "string",
		"ANONYMOUS_ID":                "string",
		"CATEGORY":                    "string",
		"CHANNEL":                     "string",
		"CONTEXT_APP_BUILD":           "string",
		"CONTEXT_APP_NAME":            "string",
		"CONTEXT_APP_NAMESPACE":       "string",
		"CONTEXT_APP_VERSION":         "string",
		"CONTEXT_DEVICE_ID":           "string",
		"CONTEXT_DEVICE_MANUFACTURER": "string",
		"CONTEXT_DEVICE_MODEL":        "string",
		"CONTEXT_DEVICE_NAME":         "string",
		"CONTEXT_IP":                  "string",
		"CONTEXT_LOCALE":              "string",
		"CONTEXT_NETWORK_CARRIER":     "string",
		"CONTEXT_REQUEST_IP":          "string",
		"CONTEXT_SCREEN_DENSITY":      "int",
		"CONTEXT_SCREEN_HEIGHT":       "int",
		"CONTEXT_SCREEN_WIDTH":        "int",
		"CONTEXT_TRAITS_ANONYMOUS_ID": "string",
		"CONTEXT_USER_AGENT":          "string",
		"EVENT":                       "string",
		"EVENT_TEXT":                  "string",
		"ID":                          "string",
		"LABEL":                       "string",
		"NAME":                        "string",
		"ORIGINAL_TIMESTAMP":          "datetime",
		"PROPERTY_1":                  "string",
		"PROPERTY_2":                  "string",
		"PROPERTY_3":                  "string",
		"PROPERTY_4":                  "string",
		"PROPERTY_5":                  "string",
		"RECEIVED_AT":                 "datetime",
		"SENT_AT":                     "datetime",
		"TIMESTAMP":                   "datetime",
		"UUID_TS":                     "datetime",
		"VALUE":                       "int",
	},
	"TRACKS": {
		"ANONYMOUS_ID":                "string",
		"CHANNEL":                     "string",
		"CONTEXT_APP_BUILD":           "string",
		"CONTEXT_APP_NAME":            "string",
		"CONTEXT_APP_NAMESPACE":       "string",
		"CONTEXT_APP_VERSION":         "string",
		"CONTEXT_DEVICE_ID":           "string",
		"CONTEXT_DEVICE_MANUFACTURER": "string",
		"CONTEXT_DEVICE_MODEL":        "string",
		"CONTEXT_DEVICE_NAME":         "string",
		"CONTEXT_IP":                  "string",
		"CONTEXT_LOCALE":              "string",
		"CONTEXT_NETWORK_CARRIER":     "string",
		"CONTEXT_REQUEST_IP":          "string",
		"CONTEXT_SCREEN_DENSITY":      "int",
		"CONTEXT_SCREEN_HEIGHT":       "int",
		"CONTEXT_SCREEN_WIDTH":        "int",
		"CONTEXT_TRAITS_ANONYMOUS_ID": "string",
		"CONTEXT_USER_AGENT":          "string",
		"EVENT":                       "string",
		"EVENT_TEXT":                  "string",
		"ID":                          "string",
		"ORIGINAL_TIMESTAMP":          "datetime",
		"RECEIVED_AT":                 "datetime",
		"SENT_AT":                     "datetime",
		"TIMESTAMP":                   "datetime",
		"UUID_TS":                     "datetime",
	},
}
var ReservedKeywordsPostgreschema = map[string]map[string]string{
	"ginkgo": {
		"_from":                       "string",
		"_join":                       "string",
		"_order":                      "string",
		"_select":                     "string",
		"_where":                      "string",
		"anonymous_id":                "string",
		"category":                    "string",
		"channel":                     "string",
		"context_app_build":           "string",
		"context_app_name":            "string",
		"context_app_namespace":       "string",
		"context_app_version":         "string",
		"context_device_id":           "string",
		"context_device_manufacturer": "string",
		"context_device_model":        "string",
		"context_device_name":         "string",
		"context_ip":                  "string",
		"context_locale":              "string",
		"context_network_carrier":     "string",
		"context_request_ip":          "string",
		"context_screen_density":      "int",
		"context_screen_height":       "int",
		"context_screen_width":        "int",
		"context_traits_anonymous_id": "string",
		"context_user_agent":          "string",
		"event":                       "string",
		"event_text":                  "string",
		"id":                          "string",
		"label":                       "string",
		"name":                        "string",
		"original_timestamp":          "datetime",
		"property_1":                  "string",
		"property_2":                  "string",
		"property_3":                  "string",
		"property_4":                  "string",
		"property_5":                  "string",
		"received_at":                 "datetime",
		"sent_at":                     "datetime",
		"timestamp":                   "datetime",
		"uuid_ts":                     "datetime",
		"value":                       "int",
	},
	"tracks": {
		"anonymous_id":                "string",
		"channel":                     "string",
		"context_app_build":           "string",
		"context_app_name":            "string",
		"context_app_namespace":       "string",
		"context_app_version":         "string",
		"context_device_id":           "string",
		"context_device_manufacturer": "string",
		"context_device_model":        "string",
		"context_device_name":         "string",
		"context_ip":                  "string",
		"context_locale":              "string",
		"context_network_carrier":     "string",
		"context_request_ip":          "string",
		"context_screen_density":      "int",
		"context_screen_height":       "int",
		"context_screen_width":        "int",
		"context_traits_anonymous_id": "string",
		"context_user_agent":          "string",
		"event":                       "string",
		"event_text":                  "string",
		"id":                          "string",
		"original_timestamp":          "datetime",
		"received_at":                 "datetime",
		"sent_at":                     "datetime",
		"timestamp":                   "datetime",
		"uuid_ts":                     "datetime",
	},
}
