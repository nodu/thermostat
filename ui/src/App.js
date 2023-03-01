import React, { useState, useEffect } from "react";

function App() {
  const [temperature, setTemperature] = useState(0);
  const [newTemperature, setNewTemperature] = useState(0);
  const [newVal, setNewVal] = useState(0);

  useEffect(() => {
    handleGetTemperature();
  }, []);

  const handleGetTemperature = async () => {
    try {
      const response = await fetch("/temperature");
      const data = await response.json();
      setTemperature(data.value);
    } catch (error) {
      console.error(error);
    }
  };

  const handleSetTemperature = async () => {
    try {
      const response = await fetch("/temperature", {
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
    let position

    switch (parseInt(newVal)) {
      case 0:
        position = -.8
        break;
      case 70:
        position = .3
        break;
      case 72:
        position = .45
        break;
      case 75:
        position = .6
        break;
      default:
        break;
    }

    try {
      const response = await fetch("/temperature", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ value: position }),
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
  // {/* <button onClick={handleGetTemperature}>Get Temperature</button> */}
  // {/* <br /> */}

  return (
    <div>
      <hr />
      <button className="button" value={0} onClick={setSetNewTemp}>off</button>
      <button className="button" value={70} onClick={setSetNewTemp}>70</button>
      <button className="button" value={72} onClick={setSetNewTemp}>72</button>
      <button className="button" value={75} onClick={setSetNewTemp}>75</button>
      <br />
      <button className="button" onClick={handleSetTemperatureManual}>Set Temperature Manual</button>

    </div>
  );
}

export default App;
