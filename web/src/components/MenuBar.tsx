import React from "react";
import MenuBarItem from './MenuBarItem';
import '../style/menubar.css';
import redbean from '../redbean.png';

function MenuBar() {
    return (
        <div className="menu-bar-wrapper">
            <div className="menu-bar-title">Red Bean</div>
            <div style={{width: '8%'}}>
                <img style={{width: '100%'}} src={redbean} alt="" />
            </div>
            <MenuBarItem text="File" />
            <MenuBarItem text="Edit" />
            <MenuBarItem text="Image" />
        </div>
    );
}

export default MenuBar;
