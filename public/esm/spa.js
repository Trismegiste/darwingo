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
                const state = JSON.parse(event.data)
                this.statByCost = state.InfoPerCost
                this.maxes = {
                    cost: state.MaxCost,
                    count: state.MaxCount,
                    victory: state.MaxAvgVictory
                }
            }
        }
    }
})