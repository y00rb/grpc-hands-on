require '../greetpb/greet_pb'
require 'grpc'

module Greet
  # The greeting service definition.
  class Service

    include GRPC::GenericService

    self.marshal_class_method = :encode
    self.unmarshal_class_method = :decode
    self.service_name = 'greet.GreetService'

    # Sends a greeting
    rpc :Greet, GreetRequest, GreetResponse
  end

  Stub = Service.rpc_stub_class
end

def main
  stub = Greet::Stub.new('localhost:50051', :this_channel_is_insecure)
  req = Greet::GreetRequest.new
  p req
  message = stub.greet(Greet::GreetRequest.new(Greeting: Greet::Greeting.new(first_name: 'y00rb'))).result
  p "Greeting: #{message}"
end

main

