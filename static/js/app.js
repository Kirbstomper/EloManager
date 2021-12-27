document.getElementById("add-player-btn").addEventListener("click",(evt)=>{
    const playerName = document.getElementById("add-player-name").value;
    const playerElo = document.getElementById("add-player-elo").value;

    console.log("Add Player Button Clicked!!!" +"Name: " +playerName +" elo: " + playerElo);

    addPlayer(playerName, playerElo)
})
document.getElementById("update-player-btn").addEventListener("click",(evt)=>{
    const playerName = document.getElementById("add-player-name").value;
    const playerElo = document.getElementById("add-player-elo").value;

    console.log("Update Player Button Clicked!!!" +"Name: " +playerName +" elo: " + playerElo);

    updatePlayer(playerName, playerElo)
})
document.getElementById("get-player-btn").addEventListener("click",(evt)=>{
    const playerName = document.getElementById("get-player-name").value;

    console.log("Get Player Button Clicked!!!" +"Name: " + playerName);

    getPlayer(playerName)
    .then((player) =>{ document.getElementById("get-player-result").value = JSON.stringify(player)})
    
})