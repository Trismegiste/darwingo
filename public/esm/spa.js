export default (host) => ({
    poolSize: 1000,
    maxRound: 10,
    epoch: 10,
    socket: null,
    populationGraph: {},

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
                this.populationGraph = JSON.parse(event.data)
                console.log(this.populationGraph)
            }
        }
    },

    getBoundingBox() {
        let minX = Infinity
        let minY = Infinity
        let maxX = -Infinity
        let maxY = -Infinity
        for (let [cost, pop] of Object.entries(this.populationGraph)) {
            cost = parseInt(cost, 10)
            if (pop > maxY) {
                maxY = pop
            }
            if (cost > maxX) {
                maxX = cost
            }
            if (pop < minY) {
                minY = pop
            }
            if (cost < minX) {
                minX = cost
            }
        }
        const width = maxX - minX
        const height = maxY - minY

        return `${minX} ${minY} ${width} ${height}`
    }
})