function addPlayer(playerName, playerElo){
    return fetch("/add",{
        method: "POST",
        body: JSON.stringify([{
            "Tag": playerName,
            "Elo": parseInt(playerElo)
        }])
    })
}
