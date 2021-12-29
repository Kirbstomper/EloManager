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
    })
    
}

function decideMatch(playerA, playerB, result){
    return fetch("/decide",{
        method: "POST",
        body: JSON.stringify({
            "playerA": playerA,
            "playerB":playerB,
            "result": result
        })
    })
}

function getPlayerList(){
    return fetch("/getAllPlayers",{
        method: "GET",
    })
}
