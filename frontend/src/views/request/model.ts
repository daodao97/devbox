export interface API {
    id: Number
    gid: Number
    method: String,
    name: String,
    url: String,
    paramsSource?: String,
    headerSource?: String,
    bodySource?: String,
    script?: String,
}

export interface Group {
    id: Number
    name: String
    children: Array<API>
}