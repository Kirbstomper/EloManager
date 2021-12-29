document.getElementById("add-player-btn").addEventListener("click",(evt)=>{
    const playerName = document.getElementById("add-player-name").value;
    const playerElo = document.getElementById("add-player-elo").value;
    if(validateIsBlank(playerName, "Player Tag")){
        return
    }
    if(validateIsBlank(playerElo, "Player Elo")){
        return
    }
    if(!parseInt(playerElo)){
        updateNotification("Elo must be a number")
        return
    }
    if(parseInt(playerElo) < 0){
        updateNotification("Elo cannot be negative")
        return
    }
    console.log("Add Player Button Clicked!!!" +"Name: " +playerName +" elo: " + playerElo);

    addPlayer(playerName, playerElo)
    .then( async (response) => {
        if (response.ok){
            let data = await response.json();
            updateNotification(JSON.stringify(data));
        }else{
            let data = await response.text();
            updateNotification(data);
        }
    })
})
document.getElementById("update-player-btn").addEventListener("click",(evt)=>{
    const playerName = document.getElementById("add-player-name").value;
    const playerElo = document.getElementById("add-player-elo").value;
    if(validateIsBlank(playerName, "Player Tag")){
        return
    }
    if(validateIsBlank(playerElo, "Player Elo")){
        return
    }
    if(!parseInt(playerElo)){
        updateNotification("Elo must be a number")
        return
    }
    if(parseInt(playerElo) < 0){
        updateNotification("Elo cannot be negative")
        return
    }
    console.log("Update Player Button Clicked!!!" +"Name: " +playerName +" elo: " + playerElo);

    updatePlayer(playerName, playerElo)
    .then( async (response) => {
        if (response.ok){
            let data = await response.json();
            updateNotification(JSON.stringify(data));
        }else{
            let data = await response.text();
            updateNotification(data);
        }
    })
})
document.getElementById("get-player-btn").addEventListener("click",(evt)=>{
    const playerName = document.getElementById("get-player-name").value;
    if(validateIsBlank(playerName, "Player Tag")){
        return
    }
    console.log("Get Player Button Clicked!!!" +"Name: " + playerName);
    getPlayer(playerName)
    .then( async (response) => {
        if (response.ok){
            let data = await response.json();
            updateNotification(JSON.stringify(data));
        }else{
            let data = await response.text();
            updateNotification(data);
        }
    })
})
document.getElementById("decide-btn").addEventListener("click",(evt)=>{
    const playerNameA = document.getElementById("decide-player-a").value;
    const playerNameB = document.getElementById("decide-player-b").value;
    const result = document.getElementById("decide-result").value;

    if(playerNameA === playerNameB){
        updateNotification("Player A and B cannot be the same!")
        return
    }
    if(validateIsBlank(playerNameA, "Player A Tag")){
        return
    }
    if(validateIsBlank(playerNameB, "Player B Tag")){
        return
    }

    console.log("Decide Match Button Clicked!!!" +"NameA: " + playerNameA + " NameB" + playerNameB + " Result:" + result);
    decideMatch(playerNameA, playerNameB, result)
    .then( async (response) => {
        if (response.ok){
            let data = await response.json();
            updateNotification(JSON.stringify(data));
        }else{
            let data = await response.text();
            updateNotification(data);
        }
    })
})

document.getElementById("get-player-list-btn").addEventListener("click", (evt)=>{
    console.log("Get Player List button clicked!")
    getPlayerList()
    .then( async (response) => {
        if (response.ok){
            let data = await response.json();
            updateNotification(JSON.stringify(data))
        }
    })
})

function updateNotification(value){
    const noti = document.getElementById("alert-box");

    noti.innerHTML = value;
    noti.parentElement.style.display = 'block';
}

function validateIsBlank(field, fieldName){
    if(field === ""){
        updateNotification(fieldName + " cannot be blank!")
        return true
    }
    return false
}