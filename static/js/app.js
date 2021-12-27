document.getElementById("addPlayer").addEventListener("click",(evt)=>{
    const playerName = document.getElementById("add-player-name").value;
    const playerElo = document.getElementById("add-player-elo").value;

    console.log("Add Player Button Clicked!!!" +"Name: " +playerName +" elo: " + playerElo);

    addPlayer(playerName, playerElo)
})