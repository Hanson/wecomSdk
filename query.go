package wecom

import (
    "net/url"
    "reflect"
    "strconv"
    "strings"
)

func toQueryValues(req any) url.Values {
    if req == nil { return url.Values{} }
    switch v := req.(type) {
    case url.Values:
        return v
    case map[string]string:
        vals := url.Values{}
        for k, s := range v { vals.Set(k, s) }
        return vals
    case map[string][]string:
        vals := url.Values{}
        for k, arr := range v { for _, s := range arr { vals.Add(k, s) } }
        return vals
    }
    rv := reflect.ValueOf(req)
    if rv.Kind() == reflect.Pointer { rv = rv.Elem() }
    vals := url.Values{}
    if rv.Kind() != reflect.Struct { return vals }
    rt := rv.Type()
    for i := 0; i < rt.NumField(); i++ {
        f := rt.Field(i)
        name := f.Tag.Get("json")
        if name == "" { name = strings.ToLower(f.Name) }
        if idx := strings.Index(name, ","); idx >= 0 { name = name[:idx] }
        if name == "-" || name == "" { continue }
        fv := rv.Field(i)
        if !fv.IsValid() { continue }
        switch fv.Kind() {
        case reflect.String:
            s := fv.String(); if s != "" { vals.Set(name, s) }
        case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
            vals.Set(name, strconv.FormatInt(fv.Int(), 10))
        case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
            vals.Set(name, strconv.FormatUint(fv.Uint(), 10))
        case reflect.Bool:
            if fv.Bool() { vals.Set(name, "1") } else { vals.Set(name, "0") }
        case reflect.Slice:
            switch fv.Type().Elem().Kind() {
            case reflect.String:
                arr := make([]string, fv.Len())
                for j := 0; j < fv.Len(); j++ { arr[j] = fv.Index(j).String() }
                vals.Set(name, strings.Join(arr, ","))
            case reflect.Int, reflect.Int64, reflect.Int32:
                arr := make([]string, fv.Len())
                for j := 0; j < fv.Len(); j++ { arr[j] = strconv.FormatInt(fv.Index(j).Int(), 10) }
                vals.Set(name, strings.Join(arr, ","))
            }
        }
    }
    return vals
}
