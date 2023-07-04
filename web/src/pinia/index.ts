import usePiniaApp from "./modules/app";
import usePiniaAdmin from "./modules/admin";

const useStore = () => ({
    app: usePiniaApp(),
    admin: usePiniaAdmin(),
});

export default useStore;