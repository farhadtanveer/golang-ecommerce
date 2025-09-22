import React, { useState, useEffect } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { useAuth } from "../context/AuthContext";

const EditItem = () => {
  const { id } = useParams();
  const navigate = useNavigate();
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [price, setPrice] = useState("");
  const [message, setMessage] = useState(null);
  const { role, token } = useAuth();

  useEffect(() => {
    if (role !== "admin") {
      setMessage({
        type: "error",
        text: "You do not have permission to edit items.",
      });
      // Optionally redirect or disable form
      // navigate('/');
      return;
    }
    fetchProduct();
  }, [id, role, navigate, token]); // Added token to dependency array

  const fetchProduct = async () => {
    try {
      const headers = {};
      if (token) {
        headers["Authorization"] = `Bearer ${token}`;
      }
      const response = await fetch(`http://localhost:8080/products/${id}`, {
        headers,
      });
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data = await response.json();
      setName(data.name);
      setDescription(data.description);
      setPrice(data.price);
    } catch (error) {
      console.error("Error fetching product for edit:", error);
      setMessage({
        type: "error",
        text: "Failed to load product for editing.",
      });
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (role !== "admin") {
      setMessage({
        type: "error",
        text: "You do not have permission to edit items.",
      });
      return;
    }

    const updatedProduct = {
      id: parseInt(id),
      name,
      description,
      price: parseFloat(price),
    };

    try {
      const response = await fetch(`http://localhost:8080/products/${id}`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify(updatedProduct),
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      setMessage({ type: "success", text: "Product updated successfully!" });
      navigate("/"); // Redirect to home after successful update
    } catch (error) {
      console.error("Error updating product:", error);
      setMessage({ type: "error", text: "Failed to update product." });
    }
  };

  return (
    <div>
      <h1 className="text-3xl font-bold mb-6">Edit Product</h1>

      {message && (
        <div
          className={`p-3 mb-4 rounded ${
            message.type === "success"
              ? "bg-green-100 text-green-800"
              : "bg-red-100 text-red-800"
          }`}
        >
          {message.text}
        </div>
      )}

      {role === "admin" ? (
        <form
          onSubmit={handleSubmit}
          className="bg-white shadow-md rounded-lg p-6"
        >
          <div className="mb-4">
            <label
              htmlFor="name"
              className="block text-gray-700 text-sm font-bold mb-2"
            >
              Name:
            </label>
            <input
              type="text"
              id="name"
              className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              value={name}
              onChange={(e) => setName(e.target.value)}
              required
            />
          </div>
          <div className="mb-4">
            <label
              htmlFor="description"
              className="block text-gray-700 text-sm font-bold mb-2"
            >
              Description:
            </label>
            <textarea
              id="description"
              className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              value={description}
              onChange={(e) => setDescription(e.target.value)}
              required
            ></textarea>
          </div>
          <div className="mb-6">
            <label
              htmlFor="price"
              className="block text-gray-700 text-sm font-bold mb-2"
            >
              Price:
            </label>
            <input
              type="number"
              id="price"
              className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              value={price}
              onChange={(e) => setPrice(e.target.value)}
              step="0.01"
              required
            />
          </div>
          <button
            type="submit"
            className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
          >
            Update Product
          </button>
        </form>
      ) : (
        <p className="text-red-500">
          You do not have permission to edit items. Please login as an Admin.
        </p>
      )}
    </div>
  );
};

export default EditItem;
