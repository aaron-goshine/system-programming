#include <iostream>
#include <fstream>

int main(int argc, char ** argv) 
{
  std::ifstream src(argv[1], std::ios::binary);
  std::ofstream src(argv[2], std::ios::binary);
  dst << src.rdbuf();
}
