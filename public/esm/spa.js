export default (host) => ({
    poolSize: 1000,
    maxRound: 10,
    epoch: 10,
    socket: null,
    statByCost: {},
    maxes: {},

    run() {
        if (typeof(EventSource) === "undefined") {
            throw new Error("Sorry, your browser does not support server-sent events...")
        }

        let url = new URL(host)
        url.searchParams.append('poolSize', this.poolSize)
        url.searchParams.append('maxRound', this.maxRound)
        url.searchParams.append('epoch', this.epoch)

        // close previous socket, because this stop the server from running
        if (this.socket !== null) {
            this.socket.close()
        }
        // open a new one
        this.socket = new EventSource(url)
        this.socket.onmessage = (event) => {
            if (event.data == 'Done') {
                this.socket.close()
            } else {
                this.statByCost = JSON.parse(event.data)
                this.maxes = this.getBoundingBox()
                console.log(this.statByCost)
            }
        }
    },

    getBoundingBox() {
        let maxCost = -Infinity
        let maxCount = -Infinity
        let maxVictory = -Infinity

        for (let [cost, info] of Object.entries(this.statByCost)) {
            cost = parseInt(cost, 10)
            if (info.GroupCount > maxCount) {
                maxCount = info.GroupCount
            }
            if (info.AvgVictory > maxVictory) {
                maxVictory = info.AvgVictory
            }
            if (cost > maxCost) {
                maxCost = cost
            }          
        }
      
        return {cost:maxCost, count:maxCount, victory:maxVictory}
    }
})