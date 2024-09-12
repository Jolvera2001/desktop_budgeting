import { users } from 'wailsjs/go/models';
import { create } from 'zustand';

interface UserState {
    selectedUser: users.User | null;
    setUser: (user: users.User | null) => void;
}

const useUserStore = create<UserState>()((set) => ({
    selectedUser: null,
    setUser: (user) => set({ selectedUser: user}),
}));