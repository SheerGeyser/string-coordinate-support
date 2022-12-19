step = 0
arr = []

function readTextFile() {
    var rawFile = new XMLHttpRequest();
    rawFile.open("GET", "data.txt", true);
    rawFile.onreadystatechange = function() {
        if (rawFile.readyState === 4) {
            var allText = rawFile.responseText;
            i = 0
            for (const e of allText.split("\n")) {
                arr[i] = e
                i++
            }
            document.querySelectorAll(".step").forEach(e => e.innerHTML = arr[step])
            document.getElementById("startBtn").style.display = "none"
            document.getElementById("nextBtn").style.display = ""
            document.getElementById("lastBtn").style.display = ""
        }
    }
    rawFile.send();
}

function nextStep() {
    step++
    document.querySelectorAll(".step").forEach(e => e.innerHTML = arr[step])
    document.querySelector(".prevStep").innerHTML = arr[step-1]
    document.querySelector(".stepCount").innerHTML = step
    // document.getElementById("textSection").innerHTML = arr[step];
}

function lastStep() {
    step--
    document.querySelectorAll(".step").forEach(e => e.innerHTML = arr[step])
    document.querySelector(".prevStep").innerHTML = arr[step-1]
    document.querySelector(".stepCount").innerHTML = step
    // document.getElementById("textSection").innerHTML = arr[step];
}