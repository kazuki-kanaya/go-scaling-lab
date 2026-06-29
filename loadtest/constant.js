import http from "k6/http";
import { sleep } from "k6";

export const options = {
  vus: 10,
  duration: "1m",
};

export default function () {
  http.get("http://localhost:8080/cpu?ms=100");
  sleep(1);
}
