import React, { useState, useEffect } from "react";

function App() {
  const [temperature, setTemperature] = useState(0);
  const [newTemperature, setNewTemperature] = useState(0);
  const [newVal, setNewVal] = useState(0);
  const [realTemperature, setRealTemperature] = useState(0);

  useEffect(() => {
    handleGetTemperature();
    handleGetRealTemperature();
  }, []);

  const handleGetRealTemperature = async () => {
    try {
      const response = await fetch("/api/realTemperature");
      const data = await response.json();
      setRealTemperature(data.real);
    } catch (error) {
      console.error(error);
    }
  };

  const handleGetTemperature = async () => {
    try {
      const response = await fetch("/api/temperature");
      const data = await response.json();
      setTemperature(data.set);
      setNewVal(data.set);
    } catch (error) {
      console.error(error);
    }
  };

  const handleSetTemperature = async () => {
    try {
      const response = await fetch("/api/temperature", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ value: newTemperature }),
      });
      await response.json();
      setTemperature(newTemperature);
      setNewTemperature(0);
    } catch (error) {
      console.error(error);
    }
  };

  const handleSetTemperatureManual = async () => {
    try {
      const response = await fetch("/temperature", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ value: parseInt(newVal) }),
      });
      await response.json();
      setTemperature(newVal);
      setNewTemperature(0);
    } catch (error) {
      console.error(error);
    }
  };

  const handleNewTemperatureChange = (event) => {
    setNewTemperature(parseFloat(event.target.value));
  };

  const setSetNewTemp = (e) => {
    setNewVal(e.target.value);
  };
  // {/* <input type="number" value={newTemperature} onChange={handleNewTemperatureChange} /> */}
  // {/* <button onClick={handleSetTemperature}>Set Temperature</button> */}
  // {/* <br /> */}

  return (
    <div>
      <hr />
      <p>"Set" Temperature: {newVal}</p>
      <p>"Real" Temperature: {realTemperature}</p>
      <button className="button" value={0} onClick={setSetNewTemp}>off</button>
      <button className="button" value={70} onClick={setSetNewTemp}>70</button>
      <button className="button" value={72} onClick={setSetNewTemp}>72</button>
      <button className="button" value={75} onClick={setSetNewTemp}>75</button>
      <br />
      <button className="button" onClick={handleSetTemperatureManual}>Set Temperature Manual</button>
      <button onClick={handleGetTemperature}>Get Temperature</button>

    </div>
  );
}

export default App;
