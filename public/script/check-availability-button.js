const el = document.getElementById("check-availability");
let values
if (el) {
    el.addEventListener("click", function() {
        let html = `
            <form id="check-availability-form" action="" method="post" novalidate class="needs-validation">
                <div class="form-row">
                    <div class="col">
                        <div class="form-row" id="reservation-dates-modal">
                            <div class="col mb-3">
                                <input disabled required class="form-control" type="text" id="start" name="start" placeholder="Arrival"/>
                            </div>

                            <div class="col">
                                <input disabled required class="form-control" type="text" id="end" name="end" placeholder="Departure"/>
                            </div>
                        </div>
                    </div>
                </div>
            </form>
            `;
            let values = attention.custom({
                msg: html,
                title: "Choose your dates",

                willOpen: () => {
                    const elem = document.getElementById("reservation-dates-modal");
                    const rp = new DateRangePicker(elem, {
                        format: "yyyy-mm-dd",
                        showOnFocus: true,
                    });
                },

                didOpen: () => {
                    document.getElementById("start").removeAttribute("disabled");
                    document.getElementById("end").removeAttribute("disabled");
                },

                callback: function(result) {
                    console.log("called");

                    let form = document.getElementById("check-availability-form");
                    let formData = new FormData(form);
                    formData.append("csrf_token", "{{.CSRFToken}}");

                    fetch("/search-availability-json", {
                        method: 'POST',
                        body: formData,
                    })
                        .then(response => response.json())
                        .then(data => {
                            console.log(data);
                            console.log(data.ok);
                            console.log(data.message);
                        })
                }
            })
            console.log(values);
        // console.log(document.getElementById("start").value,
        //     document.getElementById("end").value);
    });
}
