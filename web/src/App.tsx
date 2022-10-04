import React from 'react';
import redbean from './redbean.png';
import './App.css';
import './style/main.css';
import MenuBar from './components/MenuBar';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <MenuBar />

        <div className='image-container'>
          <img src={redbean} alt="image-container" />
        </div>
        
      </header>
    </div>
  );
}

export default App;
