function historyToBoss(data,info){
    let BossName = info.name;
    let BossHP = info.HP;
    if(data == undefined || data == null){
        return [
            {
              "round" : 1,
              "name" : BossName[0],
              "AHP" : BossHP[0],
              "NHP" : BossHP[0],
              "history" : [
                {
                  "player" : "player",
                  "time" : "0000-0-00 00:00",
                  "damage" : 0
                },
              ]
            }
          ]
    }
    let round, boss = 1;
    let results = [];
    let lastRound = data[data.length - 1].round;
    let lastBoss = data[data.length - 1].boss;
    let i = 0x00;
    for(; i < data.length; i++){
        //未击杀单独计算
        if(data[i].round == round && data[i].boss == boss){
            if(data[i].round == lastRound && data[i].boss == lastBoss) {
                results[results.length - 1].NHP += data[i].dmg;
            }
            results[results.length - 1].history.unshift({
                "player": data[i].player,
                "time": data[i].time,
                "damage": data[i].dmg
            });
        }else{
            round = data[i].round;
            boss = data[i].boss;
            results.push({
                "round": round,
                "name": BossName[boss - 1],
                "AHP": BossHP[boss - 1],
                "NHP": 0,
                "history": [{
                    "player": data[i].player,
                    "time": data[i].time,
                    "damage": data[i].dmg
                }]
            });
        }
    }
    //转换剩余血量
    results[results.length - 1].NHP = BossHP[lastBoss - 1] - results[results.length - 1].NHP
    results.reverse();
    return results;
}
export {
    historyToBoss
}