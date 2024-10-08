export default (host) => ({
    poolSize: 1000,
    maxRound: 20,
    maxEpoch: 10,
    socket: null,
    statByCost: {},
    maxes: {
        cost: 30
    },
    currentEpoch: null,
    infoBox: null,

    run() {
        if (typeof(EventSource) === "undefined") {
            throw new Error("Sorry, your browser does not support server-sent events...")
        }

        let url = new URL(host)
        url.searchParams.append('poolSize', this.poolSize)
        url.searchParams.append('maxRound', this.maxRound)
        url.searchParams.append('epoch', this.maxEpoch)

        // close previous socket, because this stop the server from running
        if (this.socket !== null) {
            this.socket.close()
        }
        // open a new one
        this.socket = new EventSource(url)
        this.socket.onmessage = (event) => {
            if (event.data == 'Done') {
                // the simulation has ended and the socket is closed on the server-side.
                // We close the socket on the client-side to prevent re-openning, which restarts the simulation
                this.socket.close()
            } else {
                const state = JSON.parse(event.data)
                this.statByCost = state.InfoPerCost
                const lastMax = this.maxes.cost // this line to prevent the graph continuously jumping, it's only growing
                this.maxes = {
                    cost: state.MaxCost > lastMax ? state.MaxCost : lastMax,
                    count: state.MaxCount,
                    victory: state.MaxVictory
                }
                this.currentEpoch = state.Epoch
            }
        }
    },

    popover(event) {
        this.infoBox = event.currentTarget.dataset.detail
    },

    popout(event) {
        this.infoBox = null
    }

})