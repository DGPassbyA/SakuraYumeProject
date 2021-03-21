function historyToBoss(data){
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