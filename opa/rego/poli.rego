package auth

default allow = false

allow {
	permit
}

permit = path {
	some i, j, k
    data.role[i].project == claims.project
    data.role[i].permission[j].role == claims.roles[_]
    
    # 遇到匹配{id}的问题
    data.role[i].permission[j].source[k].method == upper(input.attributes.request.http.headers.method)
	path := substring(data.role[i].permission[j].source[k].path, 0, indexof(data.role[i].permission[j].source[k].path, "{") - 1)
    glob.match(concat("", [path, "*"]), ["/"], input.attributes.request.http.headers.path)
}

claims = payload {
	io.jwt.verify_hs256(bearer_token, "JleHAiOjE1MDAwLCJpc3M")
	[_, payload, _] := io.jwt.decode(bearer_token)
}

bearer_token := t {
	v := input.attributes.request.http.headers.authorization
	startswith(v, "Bearer ")
	t := substring(v, count("Bearer "), -1)
}

