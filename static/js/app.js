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