import HttpService from "./HttpService";
const VehiculosService = {
    async agregarVehiculo(vehiculo) {
        return await HttpService.post("/agregar-embarazada", vehiculo);
    },
    async obtenerVehiculos() {
        return await HttpService.get("/ListaEmbarazada");
    },
    async obtenerMunicipios(){
        return await HttpService.get("/agregar-embarazada");
       
    }
};
export default VehiculosService;