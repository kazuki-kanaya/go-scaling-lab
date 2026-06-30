import http from "k6/http";
import {sleep} from "k6";

export const options = {
    stages: [
        { duration: "30s", target: 10} ,
        { duration: "1m", target: 30 },
        { duration: "30s", target: 0 },
    ]
}

export default function() {
    http.get("http://localhost/cpu?ms=500");
    sleep(1);
}
