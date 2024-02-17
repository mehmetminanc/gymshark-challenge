const API_ENDPOINT = "https://7rcijo7emg.execute-api.eu-central-1.amazonaws.com/Prod/compute-packs";

const resultElm = document.getElementById("result");
const theForm = document.getElementById("order-config");
const theButton = document.getElementById("calculate-button");

function listResult(obj) {
    resultElm.innerHTML += '<ul class="list-group">'
    for (let key in obj) {
        resultElm.innerHTML += `<li class="list-group-item">${key}: ${obj[key]}</li>`
    }
    resultElm.innerHTML += '</ul>'
}

function setLoading() {
    theButton.disabled = true
    theButton.innerHTML = `<span class="spinner-grow spinner-grow-sm" role="status" aria-hidden="true"></span> Loading...`
}

function unsetLoading() {
    theButton.disabled = false
    theButton.innerHTML = "Calculate"

}

theForm.addEventListener("submit", e => {
    e.preventDefault();
    setLoading();

    resultElm.innerText = "";

    let form = document.getElementById("order-config");
    let formData = new FormData(form);

    let order = parseInt(formData.get("order-count"), 10)
    if (isNaN(order)) {
        unsetLoading()
        alert("Please provide order count")
        return
    }
    let sizes = formData.get("pack-sizes")
        .split(",")
        .filter(x => x.trim().length)
        .filter(x => !isNaN(x))
        .map(Number)


    fetch(API_ENDPOINT, {
        method: "POST", mode: "cors", headers: {
            "Content-Type": "application/json", "Accept": "application/json"
        }, body: JSON.stringify({
            order: order, sizes: sizes,
        })
    })
        .then(response => {
            if (response.status === 200) {
                return response.json();
            } else {
                response.text().then(err => resultElm.innerHTML = "error: " + err)
            }
        })
        .then(result => {
            listResult(result);
        })
        .finally(unsetLoading);
});


