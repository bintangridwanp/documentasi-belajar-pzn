import {EventEmitter} from "events";

const  emitter = new EventEmitter;

emitter.addListener("harus-sama", (name) => {
    console.info(`Hallo  ${name}`);
});

emitter.addListener("harus-sama", (name) => {
    console.info(`Hallo  ${name}`);
});

emitter.emit("harus-sama", "Zoro");