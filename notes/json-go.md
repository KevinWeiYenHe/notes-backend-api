# JSON Go Notes


| Go type | ⇒ | JSON type |
| -- | -- | -- |
| **bool** | ⇒ | JSON boolean |
| **string** | ⇒ | JSON string |
| **int***, **uint***, **float***, **rune*** | ⇒ | JSON number |
| **array, slice** | ⇒ | JSON array |
| **struct, map** | ⇒ | JSON object |
| **nil pointers, interface values, slices, maps, etc.** | ⇒ | JSON null |
| **chan, func, complex*** | ⇒ | Not supported |
| **time.Time** | ⇒ | RFC3339-format JSON string |
| **[]byte** | ⇒ | Base64-encoded JSON string |

---

| Go type | ⇒ | JSON type |
| -- | -- | -- |
| **bool** | ⇒ | JSON boolean |
| **string** | ⇒ | JSON string |
| **int***, **uint***, **float***, **rune** | ⇒ | JSON number |
| **array, slice** | ⇒ | JSON array |
| **struct, map** | ⇒ | JSON object |
| **nil pointers, interface values, slices, maps, etc.** | ⇒ | JSON null |
| **chan, func, complex*** | ⇒ | Not supported |
| **time.Time** | ⇒ | RFC3339-format JSON string |
| **[]byte** | ⇒ | Base64-encoded JSON string |

---
