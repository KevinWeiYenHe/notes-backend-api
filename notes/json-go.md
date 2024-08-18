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

Enveloping response data like this isn’t strictly necessary, and whether you choose to do so
is partly a matter of style and taste. But there are a few tangible benefits:

1. Including a key name (like "movie") at the top-level of the JSON helps make the response more self-documenting. For any humans who see the response out of context, it is a bit easier to understand what the data relates to.
2. It reduces the risk of errors on the client side, because it’s harder to accidentally process one response thinking that it is something different. To get at the data, a client must explicitly reference it via the "movie" key.
3. If we always envelope the data returned by our API, then we mitigate a [security vulnerability](https://haacked.com/archive/2008/11/20/anatomy-of-a-subtle-json-vulnerability.aspx/) in older browsers which can arise if you return a JSON array as a response.

---

A Specification for Building APIs in JSON
https://jsonapi.org/

https://github.com/omniti-labs/jsend

---