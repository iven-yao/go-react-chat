import React, { Component, FC } from "react";
import { Navigate, Route } from "react-router-dom";

export type ProtectedRouteProps = {
    isAuth: boolean;
    redirectPath: string;
    outlet: JSX.Element;
  };
  
  export default function ProtectedRoute({isAuth, redirectPath, outlet}: ProtectedRouteProps) {
    if(isAuth) {
      return outlet;
    } else {
      return <Navigate to={{ pathname: redirectPath }} />;
    }
  };