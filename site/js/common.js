let socketUrl = "ws://1.15.20.165:7000/ws";
let apiUrl = "http://1.15.20.165:7070";

function getLocalStorage(name) {
    let value = localStorage.getItem(name);
    if (value == null) {
        return "";
    }
    return value;
}

function setLocalStorage(name, value) {
    localStorage.setItem(name, value)
}
