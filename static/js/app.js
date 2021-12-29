document.getElementById("add-player-btn").addEventListener("click",(evt)=>{
    const playerName = document.getElementById("add-player-name").value;
    const playerElo = document.getElementById("add-player-elo").value;
    if(playerName === ""){
        updateNotification("Player Tag cannot be empty")
        return
    }
    if(playerElo === ""){
        updateNotification("Player Elo cannot be empty")
        return
    }
    if(!parseInt(playerElo)){
        updateNotification("Elo must be a number")
        return
    }
    console.log("Add Player Button Clicked!!!" +"Name: " +playerName +" elo: " + playerElo);

    addPlayer(playerName, playerElo)
})
document.getElementById("update-player-btn").addEventListener("click",(evt)=>{
    const playerName = document.getElementById("add-player-name").value;
    const playerElo = document.getElementById("add-player-elo").value;
    if(playerName === ""){
        updateNotification("Player Tag cannot be empty")
        return
    }
    if(playerElo === ""){
        updateNotification("Player Elo cannot be empty")
        return
    }
    if(!parseInt(playerElo)){
        updateNotification("Elo must be a number")
        return
    }
    console.log("Update Player Button Clicked!!!" +"Name: " +playerName +" elo: " + playerElo);

    updatePlayer(playerName, playerElo)
    
})
document.getElementById("get-player-btn").addEventListener("click",(evt)=>{
    const playerName = document.getElementById("get-player-name").value;
    if(playerName ===""){
        updateNotification("Player Tag cannot be empty")
        return 
    }
    console.log("Get Player Button Clicked!!!" +"Name: " + playerName);
    getPlayer(playerName)
    .then((player) =>{ updateNotification(JSON.stringify(player))})
})
document.getElementById("decide-btn").addEventListener("click",(evt)=>{
    const playerNameA = document.getElementById("decide-player-a").value;
    const playerNameB = document.getElementById("decide-player-b").value;
    const result = document.getElementById("decide-result").value;

    if(playerNameA === ""){
        updateNotification("Player A Tag cannot be empty!")
        return
    }
    if(playerNameB === ""){
        updateNotification("Player B Tag cannot be empty!")
    }
    
    console.log("Decide Match Button Clicked!!!" +"NameA: " + playerNameA + " NameB" + playerNameB + " Result:" + result);
    decideMatch(playerNameA, playerNameB, result)
})

function updateNotification(value){
    const noti = document.getElementById("alert-box");

    noti.innerHTML = value;
    noti.parentElement.style.display = 'block';
}