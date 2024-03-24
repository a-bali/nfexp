import dayjs from "dayjs";

// format date
export function formatDate(d) {
    return dayjs(d).format("YYYY/MM/DD HH:mm:ss");
}

// convert bytes to human readable KiB/MiB/etc
export function formatBytes(bytes, decimals = 2) {
  if (!+bytes) return "0";
  const k = 1024;
  const dm = decimals < 0 ? 0 : decimals;
  const sizes = ["", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB", "ZiB", "YiB"];
  const i = Math.floor(Math.log(bytes) / Math.log(k));
  return `${parseFloat((bytes / Math.pow(k, i)).toFixed(dm))} ${sizes[i]}`;
}

// return url for api call
export function apiUrl(u) {
  return (import.meta.env.VITE_API_BASEURL || "") + "/api/" + u;
}

// form body for POST request
export function postBody(params) {
  return {
    method: "POST",
    headers: { "Content-Type": "application/x-www-form-urlencoded" },
    body: new URLSearchParams(params),
  };
}

// convert CSV to array/map format
export function csv2map(csv) {
  var map = [];
  var header = csv.split("\n")[0].split(",");

  const headernames = {
    ts: "ts",
    te: "te",
    td: "td",
    pr: "pr",
    val: "val",
    fl: "fl",
    flP: "flP",
    pkt: "pkt",
    pktP: "pktP",
    byt: "byt",
    bytP: "bytP",
    pps: "pps",
    bps: "bps",
    bpp: "bpp",
  };

  csv
    .split("\n")
    .slice(1)
    .forEach((line) => {
      if (line.length > 0) {
        var l = {};
        line.split(",").forEach((e, i) => {
          l[headernames[header[i]]] = e;
        });
        map.push(l);
      }
    });
  return map;
}

