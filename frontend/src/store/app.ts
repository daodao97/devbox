import { defineStore } from 'pinia'

export const useAppStore = defineStore({
  id: 'app',
  state: () => ({
    activeWsId: 1
  }),
  actions: {
    SetActiveWs(id: number) {
      this.activeWsId = id
    }
  }
})
