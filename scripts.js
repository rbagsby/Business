//Button for Vetspertise
document.getElementById('quick-view-btn').onclick = function() {
    document.getElementById('quick-view-popup').style.display = 'block';
}

document.getElementById('close-btn').onclick = function() {
    document.getElementById('quick-view-popup').style.display = 'none';
}

window.onclick = function(event) {
    if (event.target == document.getElementById('quick-view-popup')) {
        document.getElementById('quick-view-popup').style.display = 'none';
    }
}



//Button for Vet Peeves
document.getElementById('quick-view-btn1').onclick = function() {
    document.getElementById('quick-view-popup1').style.display = 'block';
}

document.getElementById('close-btn1').onclick = function() {
    document.getElementById('quick-view-popup1').style.display = 'none';
}

window.onclick = function(event) {
    if (event.target == document.getElementById('quick-view-popup1')) {
        document.getElementById('quick-view-popup1').style.display = 'none';
    }
}



//Button for Vetspertise Tech edition
document.getElementById('quick-view-btn2').onclick = function() {
    document.getElementById('quick-view-popup2').style.display = 'block';
}

document.getElementById('close-btn2').onclick = function() {
    document.getElementById('quick-view-popup2').style.display = 'none';
}

window.onclick = function(event) {
    if (event.target == document.getElementById('quick-view-popup2')) {
        document.getElementById('quick-view-popup2').style.display = 'none';
    }
}