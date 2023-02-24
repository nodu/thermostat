import React, { useState, useEffect } from "react";

function App() {
  const [temperature, setTemperature] = useState(0);
  const [newTemperature, setNewTemperature] = useState(0);

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

  const handleNewTemperatureChange = (event) => {
    setNewTemperature(parseFloat(event.target.value));
  };

  return (
    <div>
      <h1>Current Temperature: {temperature}</h1>
      <input type="number" value={newTemperature} onChange={handleNewTemperatureChange} />
      <button onClick={handleSetTemperature}>Set Temperature</button>
      <button onClick={handleGetTemperature}>Get Temperature</button>
    </div>
  );
}

export default App;
