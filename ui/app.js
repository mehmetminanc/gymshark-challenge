const resultElm = document.getElementById("result");
const API_ENDPOINT = "https://7rcijo7emg.execute-api.eu-central-1.amazonaws.com/Prod/compute-packs";
const theForm = document.getElementById("order-config");

theForm.addEventListener("submit", e => {
    e.preventDefault();

    resultElm.innerText = "";

    let form = document.getElementById("order-config");
    let formData = new FormData(form);

    let order = parseInt(formData.get("order-count"), 10)
    if (isNaN(order)) {
        alert("Please provide order count")
        return
    }
    let sizes = formData.get("pack-sizes")
        .split(",")
        .filter(x => x.trim().length)
        .filter(x => !isNaN(x))
        .map(Number)


    fetch(API_ENDPOINT, {
        method: "POST",
        mode: "cors",
        headers: {
            "Content-Type": "application/json",
            "Accept": "application/json"
        },
        body: JSON.stringify({
            order: order,
            sizes: sizes,
        })
    })
        .then(response => {
            return response.json();
        })
        .then(obj => {

            resultElm.innerHTML += '<ul class="list-group">'
            for (let key in obj) {
                resultElm.innerHTML += `<li class="list-group-item">${key}: ${obj[key]}</li>`
            }
            resultElm.innerHTML += '</ul>'
        });

});


