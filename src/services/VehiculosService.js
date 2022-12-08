import HttpService from "./HttpService";
const VehiculosService = {
    async agregarVehiculo(vehiculo) {
        return await HttpService.post("/agregar-embarazada", vehiculo);
    },
    async obtenerVehiculos() {
        return await HttpService.get(`/ListaEmbarazada`);
    },
};
export default VehiculosService;