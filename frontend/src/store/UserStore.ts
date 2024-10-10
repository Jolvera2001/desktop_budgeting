import { models } from 'wailsjs/go/models';
import { create } from 'zustand';

interface UserState {
    selectedUser: models.User | null;
    setUser: (user: models.User | null) => void;
}

const useUserStore = create<UserState>()((set) => ({
    selectedUser: null,
    setUser: (user) => set({ selectedUser: user}),
}));

export default useUserStore