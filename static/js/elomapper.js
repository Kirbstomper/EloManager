function addPlayer(playerName, playerElo){
    return fetch("/add",{
        method: "POST",
        body: JSON.stringify([{
            "Tag": playerName,
            "Elo": parseInt(playerElo)
        }])
    })
}
function updatePlayer(playerName, playerElo){
    return fetch("/updateElo",{
        method: "POST",
        body: JSON.stringify({
            "Tag": playerName,
            "Elo": parseInt(playerElo)
        })
    })
}

function getPlayer(playerName){
    return fetch("/getPlayer",{
        method: "POST",
        body: JSON.stringify({
            "tag": playerName
        })
    }).then((response) =>response.json())
}