namespace IronfleetTestDriver
{
    using System;
    using System.Linq;
    using System.Threading;
    using System.IO;
    using System.Net;
    using System.Collections.Generic;

    class Program
    {

        static void usage()
        {
            // Console.WriteLine("Expected usage: clientIP IP0 port0 IP1 port1 IP2 port2 num_threads duration_secs [send_reqs_at_once]");
            Console.WriteLine("Expected usage: clientIP IP_1 port_1 ... IP_n port_n num_threads duration_secs output_dir");
        }

        static void Main(string[] args)
        {            
            if (args.Length < 10)  // length 10 gives n=3 nodes, the paxos minimum for f=1
            {
                usage();
                return;
            }
            
            string guid = Guid.NewGuid().ToString();

            ulong num_threads = 1;
            ulong experiment_duration = 60;
            string output_directory = String.Format("IronfleetOutput/Job-{0}", guid);
            IPAddress client_ip;
            bool send_reqs_at_once = false;
            
            ClientBase.endpoints = new List<IPEndPoint>();

            try
            {
                output_directory = args[args.Length-1];
                experiment_duration = Convert.ToUInt64(args[args.Length-2]);
                num_threads = Convert.ToUInt64(args[args.Length-3]);
                client_ip = IPAddress.Parse(args[0]);

                for (int i = 1; i < args.Length-3; i=i+2) {
                    ClientBase.endpoints.Add(new IPEndPoint(IPAddress.Parse(args[i]), Convert.ToInt32(args[i+1])));
                // TONY: Delete this functionality for now
                // if (args.Length > 9)
                // {
                //     send_reqs_at_once = true;
                // }
                }
            }
            catch (Exception e)
            {
                Console.WriteLine("Command line exception: " + e);
                usage();
                return;
            }
            ClientBase.my_addr = client_ip;

            // Create a directory for logging all of our output
            // string output_directory = String.Format("{0}\\IronfleetOutput\\Job-{1}",
            //     System.Environment.GetEnvironmentVariable("TMP"),
            //     guid);
            Directory.CreateDirectory(output_directory);
            Multipaxos.Client.Trace("Output directory " + output_directory);

            // Create the log file itself
            // FileStream log = new FileStream(output_directory + "\\client.txt", FileMode.Create);
            FileStream log = new FileStream(output_directory + "/client.txt", FileMode.Create);
            StreamWriter log_stream = new StreamWriter(log);

            HiResTimer.Initialize();
            Multipaxos.Client.Trace("Client process starting " + num_threads + "thread(s) running for "+ experiment_duration + "s ...");
            string targets = "";
            foreach (IPEndPoint i in ClientBase.endpoints) {
                targets += i + ", ";
            }
            Multipaxos.Client.Trace("Targets are " + targets);
            
            Console.WriteLine("[[READY]]");
            Console.WriteLine("ClientGUID {0}", guid);
            
            // Redirect all subsequent output to the log
            TextWriter stdout = Console.Out;
            Console.SetOut(log_stream);

            // Start the experiment
            var threads = ClientBase.StartThreads<Multipaxos.Client>(num_threads, send_reqs_at_once).ToArray();

            if (experiment_duration == 0)
            {
                threads[0].Join();
            }
            else
            {
                Thread.Sleep((int)experiment_duration * 1000);
                stdout.WriteLine("[[DONE]]");
                stdout.Flush();
                log_stream.Flush();
                Environment.Exit(0);
            }
        }
    }
}
