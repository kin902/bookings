const el = document.getElementById("check-availability");

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
                                <input disabled required class="form-control" type="text" id="end" name="ned" placeholder="Departure"/>
                            </div>
                        </div>
                    </div>
                </div>
            </form>
            `;
        attention.custom({msg: html, title: "Choose your dates"})
    });
}