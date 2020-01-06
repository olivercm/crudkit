syntax = "proto3";

package service.{{.Server}};

service {{.UCFirstServer}}Service {
    rpc Get{{.Model}}List (Get{{.Model}}ListReq) returns (Get{{.Model}}ListResp);
    rpc Create{{.Model}} (Create{{.Model}}Req) returns (Create{{.Model}}Resp);
    rpc Update{{.Model}} (Update{{.Model}}Req) returns (Update{{.Model}}Resp);
    rpc Delete{{.Model}} (Delete{{.Model}}Req) returns (Delete{{.Model}}Resp);
}

message Get{{.Model}}ListReq {
    int64 currentPage = 1;
    int64 pageSize = 2;
}

message Get{{.Model}}ListResp {
    repeated {{.Model}}ListData result = 1;
}

message {{.Model}}ListData {
    {{range .ProtoFields}}{{.}}
    {{end}}
}

message Create{{.Model}}Req {
    {{range .ProtoFields}}{{.}}
    {{end}}
}

message Create{{.Model}}Resp {

}

message Update{{.Model}}Req {
    {{range .ProtoFields}}{{.}}
    {{end}}
}

message Update{{.Model}}Resp {

}

message Delete{{.Model}}Req {
    int64 id = 1;
}

message Delete{{.Model}}Resp {

}


