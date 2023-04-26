import React, { useState, useEffect } from "react";

function App() {
  const [temperature, setTemperature] = useState(0);
  const [realTemperature, setRealTemperature] = useState(0);
  const [isLoading, setIsLoading] = useState(false);
  const [isCronEnabled, setIsCronEnabled] = useState(false);

  useEffect(() => {
    handleGetTemperature();
    handleGetRealTemperature();
    handleGetCronEnabled();
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
    } catch (error) {
      console.error(error);
    }
  };

  const handleGetCronEnabled = async () => {
    try {
      const response = await fetch("/api/cron");
      const data = await response.json();
      setIsCronEnabled(data.cron);
    } catch (error) {
      console.error(error);
    }
  };

  const handleToggleCron = async () => {
    setIsLoading(true);
    try {
      const response = await fetch("/api/cron", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ cron: !isCronEnabled }),
      });
      await response.json();
      setIsCronEnabled(!isCronEnabled);
      setIsLoading(false);
    } catch (error) {
      console.error(error);
    }
  };

  const handleTemperatureChange = async (e) => {
    setIsLoading(true);
    try {
      const response = await fetch("/api/temperature", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ Set: parseInt(e.target.value) }),
      });
      await response.json();
      setTemperature(e.target.value);
      setIsLoading(false);
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <div>
      <div className={isLoading ? "loader loader-position" : undefined}></div>
      <p>"Set" Temperature: {temperature}</p>
      <p className="text-red-700">"Real" Temperature: {realTemperature || "..."}</p>
      <button className="button" value={0} onClick={handleTemperatureChange}>off</button>
      <button className="button" value={70} onClick={handleTemperatureChange}>70</button>
      <button className="button" value={72} onClick={handleTemperatureChange}>72</button>
      <button className="button" value={75} onClick={handleTemperatureChange}>75</button>
      <button className="button" value={80} onClick={handleTemperatureChange}>80</button>
      <hr />
      <p className="text-sky-500">Schedule Enabled: {isCronEnabled ? "Yes" : "No"}</p>
      <button className="button" onClick={handleToggleCron}>Toggle Schedule Enabled</button>
    </div >
  );
}

export default App;
