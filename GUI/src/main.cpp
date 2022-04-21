#include <iostream>

#define GLFW_INCLUDE_VULKAN
#include <GLFW/glfw3.h>

#include "RubikGUIConfig.h"

void errorCallback(int error, const char *description)
{
    std::cout << "Error occurred! Code: " << error << " Description: " << description << std::endl;
}

int main(int argc, char **argv)
{
    if (glfwInit() != GLFW_TRUE)
    {
        std::cout << "Can't initialize GLFW, exiting...";

        return (-1);
    }

    glfwSetErrorCallback(errorCallback);

    GLFWwindow *window = glfwCreateWindow(640, 480, "My Title", NULL, NULL);
    if (!window)
    {
        std::cout << "Can't create GLFW Window, exiting..." << std::endl;

        return (-1);
    }

    VkInstance instance;
    VkResult result;
    VkInstanceCreateInfo info = {};

    result = vkCreateInstance(&info, nullptr, &instance);
    if (result != VK_SUCCESS)
    {
        std::cout << "Can't create Vulkan context, exiting..." << std::endl;

        return (-1);
    }

    while (!glfwWindowShouldClose(window))
    {
        ;
    }

    glfwDestroyWindow(window);
    glfwTerminate();

    return (0);
}
