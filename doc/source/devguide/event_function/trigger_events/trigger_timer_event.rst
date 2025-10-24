Timer Event Source
==================

You can schedule a timer to invoke your code based on a fixed rate of minutes,
hours, or days or a cron expression. For details, see
`Using a Timer Trigger <https://docs.otc.t-systems.com/function-graph/umn/creating_triggers/using_a_timer_trigger.html>`__.

Timer example event
-------------------

.. code-block:: json

   {
     "version": "v1.0",
     "time": "2018-06-01T08:30:00+08:00",
     "trigger_type": "TIMER",
     "trigger_name": "Timer_001",
     "user_event": "User Event"
   }

Parameter description
---------------------

.. list-table::
   :header-rows: 1
   :widths: 20 15 30 35

   * - Parameter
     - Type
     - Example Value
     - Description
   * - version
     - String
     - V1.0
     - Event version
   * - time
     - String
     - 2018-06-01T08:30:00+08:00
     - Time when an event occurs.
   * - trigger_type
     - String
     - TIMER
     - Trigger type
   * - trigger_name
     - String
     - Timer_001
     - Trigger name
   * - user_event
     - String
     - User Event
     - Additional information of the trigger

Example
-------

.. literalinclude:: /../../samples-events/timer/handler.go
    :language: go
    :caption: :github_repo_master:`handler.go <samples-events/timer/handler.go>`
