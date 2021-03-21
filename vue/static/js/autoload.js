function historyToBoss(data){
    //{"eid":1,"player":"洛夜","time":"2021-02-05T21:38:17.953491Z","round":1,"boss":1,"dmg":3044001,"flag":0}
    /*{
          "round" : 1,
          "name" : "狮鹫",
          "AHP" : 800000,
          "NHP" : parseInt(500000*Math.random()),
          "history" : [
            {
              "player" : "player1",
              "time" : "2020-9-20 16:00",
              "damage" : parseInt(1000000*Math.random())
            },
            {
              "player" : "player2",
              "time" : "2020-9-20 17:00",
              "damage" : 1200000
            },
            {
              "player" : "player3",
              "time" : "2020-9-20 20:00",
              "damage" : 1030000
            }
          ]
        }*/
    let round, boss = 1;
    let results = [];
    const BossName = ["哥布林王", "狮鹫", "棱镜夫人", "独眼巨人", "双鱼"];
    //国服数据各阶段血量暂时一样
    const BossHP = [6000000, 8000000, 10000000, 12000000, 20000000];
    let lastRound = data[data.length - 1].round;
    let lastBoss = data[data.length - 1].boss;
    let i = 0x00;
    for(; i < data.length; i++){
        //未击杀单独计算
        if(data[i].round == lastRound && data[i].boss == lastBoss) {
            break;
        }
        if(data[i].round == round && data[i].boss == boss){
            results[results.length - 1].history.append({
                "player": data[i].player,
                "time": data[i].time,
                "damage": data[i].dmg
            });
        }else{
            round = data[i].round;
            boss = data[i].boss;
            results.append({
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
    return results;
}
export {
    historyToBoss
}